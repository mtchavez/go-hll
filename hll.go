package go_hll

import (
  "math"
)

type HllInfo struct {
  map_size float64
  k_compliment uint32
  alpha float64
  Table map[uint32]uint32
}

func Hll(std_error float64) HllInfo {
  info := HllInfo{}
  map_size := float64(1.04) / std_error
  k := math.Ceil(math.Log2(map_size * map_size))
  info.k_compliment = uint32(32 - k)

  map_size = math.Pow(2, k)
  info.map_size = map_size
  info.alpha = get_alpha(map_size)
  table := make(map[uint32]uint32)
  var i uint32 = 0
  for ; i < uint32(map_size); i++ {
    table[i] = 0
  }
  info.Table = table
  return info
}

func get_alpha(map_size float64) float64 {
  if map_size == 16 {
    return 0.673
  } else if map_size == 32 {
    return 0.697
  } else if map_size == 64 {
    return 0.709
  }
  return 0.7213 / (1.0 + 1.079 / map_size)
}

func HashCode(key string) uint32 {

  var hash, i uint32 = 0, 0
  runes := []rune(key)

  for ; i < uint32(len(key)); i++ {
    hash += uint32(runes[i])
    hash += (hash << 10)
    hash ^= (hash >> 6)
  }

  hash += (hash << 3);
  hash ^= (hash >> 11);
  hash += (hash << 15);

  return hash
}

func getRank(hash_code uint32, max uint32) uint32 {
    var r uint32 = 1
    var one uint32 = 1
    for ; ((hash_code & one) == 0 && r <= max); r++ {
      hash_code = hash_code >> 1
    }
    return r
}

func Add(info HllInfo, term string) HllInfo {
  k_compliment := info.k_compliment
  hash_code := HashCode(term)
  j := uint32(hash_code >> k_compliment)

  t := info.Table
  t[j] = uint32(math.Max(float64(t[j]), float64(getRank(hash_code, k_compliment))))
  info.Table = t
  return info
}

func Count(info HllInfo) uint32 {
  t := info.Table
  var c float64 = 0
  var i uint32 = 0

  for ; i < uint32(info.map_size); i++ {
      c += 1.0 / math.Pow(2, float64(t[i]))
    }

  e := info.alpha * info.map_size * info.map_size / c

  // Make corrections & smoothen things.
  if (e <= (5 / 2) * info.map_size) {
    var v float64 = 0
    for i = 0; i < uint32(info.map_size); i++ {
        if t[i] == 0 {
          v++
        }
    }
    if v > 0 {
      e = info.map_size * math.Log(info.map_size / v)
    }
  } else {
    pow32 := math.Pow(2, 32)
    if e > (1.0 / 30.0) * pow32 {
        e = -pow32 * math.Log(1.0 - e / pow32)
      }
  }
  return uint32(e)
}
