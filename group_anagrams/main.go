package main

import (
	"fmt"
	"sort"
)

func main() {
	inp := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(inp))
}

func groupAnagrams(strs []string) [][]string {

	grams := make(map[string][]string)

	for _, str := range strs {

		r := []rune(str)
		sort.Slice(r, func(i, j int) bool {
			return r[i] < r[j]
		})

		if group, ok := grams[string(r)]; ok {
			group = append(group, str)
			grams[string(r)] = group
		} else {
			group = []string{str}
			grams[string(r)] = group

		}

	}
	result := [][]string{}
	for _, v := range grams {
		result = append(result, v)
	}

	return result
}
