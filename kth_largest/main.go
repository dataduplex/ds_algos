package main

import "fmt"

func main() {
	//fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}

func findKthLargest(nums []int, k int) int {

	if k < 0 || k > len(nums) || len(nums) == 0 {
		return 0
	}

	k = len(nums) - k + 1

	start := 0
	end := len(nums) - 1
	result := -1
	for {
		idx := partition(nums, start, end)
		if idx+1 == k {
			result = nums[idx]
			break
		} else if idx+1 < k {
			start = idx + 1
		} else {
			end = idx - 1
		}

	}

	return result

}

func partition(nums []int, start int, end int) int {

	if start == end {
		return start
	}

	pivot := nums[end]
	b := start
	for i := start; i < end; i++ {
		if nums[i] < pivot {
			nums[i], nums[b] = nums[b], nums[i]
			b++
		}
	}
	nums[b], nums[end] = nums[end], nums[b]
	return b
}
