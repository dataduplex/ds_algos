package main

import "fmt"

func main() {
	//[5,7,7,8,8,10]

	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
}

func searchRange(nums []int, target int) []int {

	//result := []int{}

	start, end := 0, len(nums)-1
	mid := binarySearch(nums, start, end, target)
	if mid == -1 {
		return []int{-1, -1}
	}
	left, right := mid, mid
	start, end = 0, mid-1
	for {
		idx := binarySearch(nums, start, end, target)
		if idx == -1 {
			break
		}
		left = idx
		end = left - 1
	}
	start, end = mid+1, len(nums)-1
	for {
		idx := binarySearch(nums, start, end, target)
		if idx == -1 {
			break
		}
		right = idx
		start = right + 1
	}

	return []int{left, right}
}

func binarySearch(nums []int, from int, to int, target int) int {

	if from > to || from < 0 || to >= len(nums) {
		return -1
	}

	mid := from + (to-from)/2

	if nums[mid] == target {
		return mid
	} else if target < nums[mid] {
		return binarySearch(nums, from, mid-1, target)
	} else {
		return binarySearch(nums, mid+1, to, target)
	}

}
