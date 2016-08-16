package hll

import (
	"math"
)

// Hll is the hyper-log-log struct containing the lookup table
// and internal configuration
type Hll struct {
	Table   map[uint32]uint32
	mapSize float64
	comp    uint32
	alpha   float64
	stdErr  float64
}

// Initialize takes in an allowed error for the hyper-log-log algorithm
// and will return an Hll
func Initialize(err float64) (hll *Hll) {
	hll = &Hll{}
	hll.setupState(err)
	hll.setupTable()
	return
}

func (hll *Hll) setupState(err float64) {
	hll.stdErr = err
	msize := float64(1.04) / hll.stdErr
	k := math.Ceil(math.Log2(msize * msize))
	hll.comp = uint32(32 - k)
	hll.mapSize = math.Pow(2, k)
	hll.alpha = getAlpha(hll.mapSize)
}

func (hll *Hll) setupTable() {
	table := make(map[uint32]uint32)
	var i uint32
	for ; i < uint32(hll.mapSize); i++ {
		table[i] = 0
	}
	hll.Table = table
}

func getAlpha(msize float64) (alpha float64) {
	if msize == 16 {
		alpha = 0.673
	} else if msize == 32 {
		alpha = 0.697
	} else if msize == 64 {
		alpha = 0.709
	} else {
		alpha = 0.7213 / (1.0 + 1.079/msize)
	}
	return
}

func hashCode(key string) (hashed uint32) {
	var i uint32
	runes := []rune(key)

	for ; i < uint32(len(key)); i++ {
		hashed += uint32(runes[i])
		hashed += (hashed << 10)
		hashed ^= (hashed >> 6)
	}

	hashed += (hashed << 3)
	hashed ^= (hashed >> 11)
	hashed += (hashed << 15)

	return
}

func getRank(hashed uint32, max uint32) (rank uint32) {
	var one uint32 = 1
	for rank = 1; (hashed&one) == 0 && rank <= max; rank++ {
		hashed = hashed >> 1
	}
	return rank
}

// Add takes a term to add to the hyper-log-log table
// and will be hashed and incremented
func (hll *Hll) Add(term string) {
	hashed := hashCode(term)
	key := uint32(hashed >> hll.comp)

	current := float64(hll.Table[key])
	rank := float64(getRank(hashed, hll.comp))
	hll.Table[key] = uint32(math.Max(current, rank))
}

// Count will calculate the size of the terms in the hyper-log-log table
// and return the total, consider that there may be false positives.
func (hll *Hll) Count() uint32 {
	var c float64
	var i uint32

	for ; i < uint32(hll.mapSize); i++ {
		c += 1.0 / math.Pow(2, float64(hll.Table[i]))
	}

	e := hll.alpha * hll.mapSize * hll.mapSize / c

	// Make corrections
	if e <= (5/2)*hll.mapSize {
		var v float64
		for i = 0; i < uint32(hll.mapSize); i++ {
			if hll.Table[i] == 0 {
				v++
			}
		}
		if v > 0 {
			e = hll.mapSize * math.Log(hll.mapSize/v)
		}
	} else {
		pow32 := math.Pow(2, 32)
		if e > (1.0/30.0)*pow32 {
			e = -pow32 * math.Log(1.0-e/pow32)
		}
	}
	return uint32(e)
}
