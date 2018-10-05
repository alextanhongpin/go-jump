package main

import (
	"fmt"
)

func main() {
	counter := make(map[int32]int)
	var buckets int32 = 10
	add := func(buckets int32, iter int) {
		for i := 0; i < iter; i++ {
			o := JumpConsistentHash(uint64(i), buckets)
			counter[o]++
		}
	}
	add(buckets, 100)
	print(counter)

	// Adding two new buckets
	add(buckets+2, 100)
	print(counter)

	// Removing two buckets
	add(buckets-2, 100)
	print(counter)
}

func print(data map[int32]int) {
	for i := 0; i < len(data); i++ {
		fmt.Println(i, data[int32(i)])
	}
	fmt.Println()
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
