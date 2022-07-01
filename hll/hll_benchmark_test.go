package hll

import (
	"bufio"
	"os"
	"testing"
)

func BenchmarkHllNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New()
	}
}

func BenchmarkHllAdd(b *testing.B) {
	hll := New()
	fd, _ := os.Open("/usr/share/dict/words")
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	var wordCount int
	var totalWords int
	var values []string
	for scanner.Scan() {
		word := scanner.Text()
		totalWords++

		hll.Add(word)
		wordCount++
		values = append(values, word)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hll.Add(values[i%totalWords])
	}
}

func BenchmarkHllCount(b *testing.B) {
	hll := New()
	fd, _ := os.Open("/usr/share/dict/words")
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		hll.Add(scanner.Text())
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hll.Count()
	}
}
