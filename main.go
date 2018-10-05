package main

import (
	"fmt"
)

func main() {
	counter := make(map[int32]int)
	var buckets int32 = 10
	for i := 0; i < 1000000; i++ {
		o := JumpConsistentHash(uint64(i), buckets)
		counter[o]++
	}

	sorted := make([]int, len(counter))
	for k, v := range counter {
		sorted[int(k)] = v
	}
	for k, v := range sorted {
		fmt.Printf("%d => %d\n", k, v)
	}

}

func JumpConsistentHash(key uint64, numBuckets int32) int32 {
	b, j := int32(-1), int32(0)
	for j < numBuckets {
		b = j
		key = key*uint64(2862933555777941757) + 1
		j = (b + 1) * int32((1<<31)/((key>>33)+1))
	}
	return b
}
