package main

import (
	"math"
	"sort"
)

func main() {}

func merge(intervals [][]int) [][]int {
	out := [][]int{}

	// we sort the intervals based on the starts to ensure that they are always either sorted or should be merged
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// the first interval so far has nothing to be merged iwth (the first one)
	out = append(out, intervals[0])

	for _, interval := range intervals {
		start, end := interval[0], interval[1]

		// if its start is less than or eauqles the previous interval end, we have to merge them togeather, and by merging, it means change the end of the previous interval
		if start <= out[len(out)-1][1] {
			out[len(out)-1][1] = int(math.Max(float64(end), float64(out[len(out)-1][1]))) // max because this is a scenario : [1, 5], [2, 4] so if without max, we would have [1, 4] and this is wrong
		} else {
			out = append(out, interval)
		}
	}

	return out
}
