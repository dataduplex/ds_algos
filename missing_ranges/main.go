package main

import "fmt"

func main() {

}

func findMissingRanges(nums []int, lower int, upper int) []string {

	result := []string{}
	if len(nums) == 0 {
		result = append(result, printRange(lower, upper))
		return result
	}

	if lower < nums[0] {
		//result = append([]string{printRange(lower,nums[0]-1)}, result...)
		result = append(result, printRange(lower, nums[0]-1))
	}
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > 1 {
			result = append(result, printRange(nums[i-1]+1, nums[i]-1))
		}
	}

	if upper > nums[len(nums)-1] {
		result = append(result, printRange(nums[len(nums)-1]+1, upper))
	}

	return result

}

func printRange(from int, to int) string {
	if from == to {
		return fmt.Sprintf("%d", from)
	}
	return fmt.Sprintf("%d", from) + "->" + fmt.Sprintf("%d", to)
}
