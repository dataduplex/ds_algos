package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	// [[0, 30],[5, 10],[15, 20]]

	//inv := [][]int{[]int{0, 30}, []int{5, 10}, []int{15, 20}}

	inv := [][]int{[]int{7, 10}, []int{2, 4}}
	fmt.Println(minMeetingRooms(inv))

}

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

func (h IntHeap) Less(i int, j int) bool { return h[i] < h[j] }

func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minMeetingRooms(intervals [][]int) int {

	if len(intervals) == 0 {
		return 0
	}

	sort.SliceStable(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	q := &IntHeap{intervals[0][1]}
	heap.Init(q)
	maxRooms := 1

	for i := 1; i < len(intervals); i++ {
		for q.Len() > 0 && (*q)[0] <= intervals[i][0] {
			heap.Pop(q)
		}
		heap.Push(q, intervals[i][1])

		if q.Len() > maxRooms {
			maxRooms = q.Len()
		}
	}

	return maxRooms

}
