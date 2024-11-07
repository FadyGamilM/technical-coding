package main

import (
	"log"
	"math"
)

func main() {
	log.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
	log.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
	log.Println(longestOnes([]int{0, 0, 0, 1}, 4))
}

func longestOnes(nums []int, k int) int {
	flippedZeros := 0
	start := 0
	maxSoFar := math.MinInt
	end := 0
	for i, num := range nums {
		end = i
		if num == 0 {
			if flippedZeros >= k {
				for flippedZeros >= k {
					if nums[start] == 0 {
						flippedZeros--
						start++
					} else {
						start++
					}
				}
				flippedZeros++
				// lets override if we got another max
				maxSoFar = int(math.Max(float64(maxSoFar), float64(float64(end)-float64(start)+1)))
			} else {
				flippedZeros++
				// lets override if we got another max
				maxSoFar = int(math.Max(float64(maxSoFar), float64(float64(end)-float64(start)+1)))
			}
		} else {
			// lets override if we got another max
			maxSoFar = int(math.Max(float64(maxSoFar), float64(float64(end)-float64(start)+1)))
		}
	}

	return maxSoFar
}
