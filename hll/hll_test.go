package hll

import (
	"fmt"
	"math"
	"testing"
)

func TestGetAlpha(t *testing.T) {
	result := getAlpha(16)
	if result != 0.673 {
		t.Errorf("Expected alpha of 0.673 but got %f", result)
	}

	result = getAlpha(32)
	if result != 0.697 {
		t.Errorf("Expected alpha of 0.697 but got %f", result)
	}

	result = getAlpha(64)
	if result != 0.709 {
		t.Errorf("Expected alpha of 0.709 but got %f", result)
	}

	result = getAlpha(128)
	if fmt.Sprintf("%f", result) != fmt.Sprintf("%f", 0.715270) {
		t.Errorf("Expected alpha of 0.7152705 but got %f", result)
	}
}

func TestInitialize(t *testing.T) {
	hll := Initialize(0.065)
	if hll.mapSize != 256.0 {
		t.Errorf("Map size should be 256 but got %d", hll.mapSize)
	}

	if len(hll.Table) != int(hll.mapSize) {
		t.Errorf("Table size should be %d but got %d", hll.mapSize, len(hll.Table))
	}

	if hll.kComp != 24 {
		t.Errorf("K compliment should be 24 but got %d", hll.kComp)
	}

	if fmt.Sprintf("%g", hll.alpha) != fmt.Sprintf("%g", 0.7182725932495458) {
		t.Errorf("Alpha should be 0.7182725932495458 but got %g", hll.alpha)
	}
}

func TestHashCode(t *testing.T) {
	result := HashCode("apple")
	var expected uint32 = 2297466611
	if result != expected {
		t.Errorf("Hash code for 'apple' should return %d but got %d", expected, result)
	}
}

func TestGetRank(t *testing.T) {
	hll := Initialize(0.065)
	hashed := HashCode("kiwi kiwi kiwi")
	rank := getRank(hashed, hll.kComp)
	var expected uint32 = 2
	if rank != expected {
		t.Errorf("Rank of 'apple' should be %d but got %d", expected, rank)
	}

	hashed = HashCode("apple")
	rank = getRank(hashed, hll.kComp)
	expected = 1
	if rank != expected {
		t.Errorf("Rank of 'apple' should be %d but got %d", expected, rank)
	}
}

func TestAdd(t *testing.T) {
	hll := Initialize(0.065)

	hll.Add("apple")
	hashed := HashCode("apple")
	key := uint32(hashed >> hll.kComp)
	value := uint32(math.Max(float64(hll.Table[key]), float64(getRank(hashed, hll.kComp))))
	var expected uint32 = 1
	if value != expected {
		t.Errorf("Value after add should be %d but got %d", expected, value)
	}

	hll.Add("apple")
	hll.Add("apple")
	value = uint32(math.Max(float64(hll.Table[key]), float64(getRank(hashed, hll.kComp))))
	expected = 1
	if value != expected {
		t.Errorf("Value after add should be %d but got %d", expected, value)
	}

	hll.Add("kiwi kiwi kiwi")
	hashed = HashCode("kiwi kiwi kiwi")
	key = uint32(hashed >> hll.kComp)
	value = uint32(math.Max(float64(hll.Table[key]), float64(getRank(hashed, hll.kComp))))
	expected = 2
	if value != expected {
		t.Errorf("Value after add should be %d but got %d", expected, value)
	}

	hll.Add("apple banana peach wiki pear")
	hashed = HashCode("apple banana peach wiki pear apple")
	key = uint32(hashed >> hll.kComp)
	value = uint32(math.Max(float64(hll.Table[key]), float64(getRank(hashed, hll.kComp))))
	expected = 4
	if value != expected {
		t.Errorf("Value after add should be %d but got %d", expected, value)
	}
}

func TestCount(t *testing.T) {
	hll := Initialize(0.065)
	words := []string{"apple", "this is bananas", "kiwi kiwi kiwi", "Peach is a peach", "apple banana peach wiki pear"}
	for _, word := range words {
		hll.Add(word)
	}
	count := hll.Count()
	if count != uint32(len(words)) {
		t.Errorf("Count should be %d but got %d", len(words), count)
	}
}
