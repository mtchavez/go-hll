package hll

import (
	"fmt"
)

func ExampleHll_Count() {
	hll := NewWithErr(0.065)
	words := []string{"apple", "this is bananas", "kiwi kiwi kiwi", "Peach is a peach", "apple banana peach wiki pear"}
	for _, word := range words {
		hll.Add(word)
	}
	count := hll.Count()
	err := float32((count - uint32(len(words))) / uint32(len(words)) / 100.0)
	fmt.Printf("\nCount: %d\nError: %f%%\n\n", count, err)
	// Output:
	//
	// Count: 5
	// Error: 0.000000%
}
