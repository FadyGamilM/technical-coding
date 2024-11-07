package main

import "log"

func main() {
	log.Println(twoSum([]int{1, 1}, 2))
	log.Println(twoSum([]int{2, 5, 7, 4}, 9))
	log.Println(twoSum([]int{3, 3}, 6))
}

func twoSum(nums []int, target int) []int {
	out := []int{}
	tracker := map[int]int{}

	for currIdx, num := range nums {
		if idx, ok := tracker[num]; ok {
			out = append(out, idx)
			out = append(out, currIdx)
			return out
		}

		tracker[target-num] = currIdx
	}

	return out

}
