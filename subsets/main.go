package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%+v\n", subsets([]int{15, 20, 12, 19, 4}))
}


func subsets(A []int) [][]int {

	sort.Ints(A)

	if len(A) == 0 {
		return [][]int{[]int{}}
	}


	temp1 := subsets(A[1:])
	temp2 := make([][]int, len(temp1))
	copy(temp2,temp1)


	for i:=0; i<len(temp2); i++ {
		temp2[i] = append([]int{A[0]}, temp2[i]...)
	}

	result := [][]int{}

	result = append(result, temp1[0])
	result = append(result, temp2...)
	result = append(result, temp1[1:]...)

	return result

}