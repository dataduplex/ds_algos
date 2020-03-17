package main

func main() {
	lengthOfLongestSubstring("abcabcbb")
}

func lengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}

	idxMap := make(map[string]int)
	start, max := 0, 0

	for i, c := range s {

		ch := string(c)

		idx, ok := idxMap[ch]
		if ok && idx >= start {
			start = idx + 1
		}

		idxMap[ch] = i
		length := i - start + 1
		//fmt.Printf("i=%d, ch=%s, len=%d\n", i, ch, length)
		if length > max {
			max = length
		}

	}

	return max
}
