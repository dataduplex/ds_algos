package main

import (
	"fmt"
	"sort"
)

func main() {
	result := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Printf("result: %+v\n", result)
}

func threeSum(nums []int) [][]int {

	sort.Ints(nums)

	/*
	   -4 -1 -1 0 1 2
	*/
	result := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		//twoSumPairs := [][]int{}
		twoSumPairs := twoSum(nums[i+1:], 0-nums[i])
		for _, pair := range twoSumPairs {
			pair = append(pair, nums[i])
			result = append(result, pair)
		}
	}

	return result
}

func twoSum(nums []int, target int) [][]int {

	/*
	   1 2 2 3 3 4
	*/

	i := 0
	j := len(nums) - 1
	result := [][]int{}
	for i < j {

		if i > 0 && nums[i] == nums[i-1] {
			i++
			continue
		}

		if j < len(nums)-1 && nums[j] == nums[j+1] {
			j--
			continue
		}

		if nums[i]+nums[j] == target {
			result = append(result, []int{nums[i], nums[j]})
			i++
			j--
		} else if nums[i]+nums[j] < target {
			i++
		} else {
			j--
		}

	}
	return result
}
