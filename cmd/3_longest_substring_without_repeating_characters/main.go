package main

import (
	"log"
	"math"
)

func main() {
	log.Println(lengthOfLongestSubstring("abcabcbb"))
	log.Println(lengthOfLongestSubstring("abcbbcbb"))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	msf := math.MinInt
	seen := map[rune]struct{}{}
	start := 0
	for end, r := range s {
		if _, seenBefore := seen[r]; seenBefore {
			for seenBefore {
				delete(seen, rune(s[start]))
				start++
				_, seenBefore = seen[r]
				if !seenBefore {
					msf = int(math.Max(float64(msf), float64(end-start+1)))
					seen[r] = struct{}{}
					break
				}
			}
		} else {
			msf = int(math.Max(float64(msf), float64(end-start+1)))
			seen[r] = struct{}{}
		}
	}

	return msf
}
