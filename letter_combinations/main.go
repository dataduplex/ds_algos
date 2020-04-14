package main

import "fmt"

func main () {
	fmt.Printf("%+v\n", letterCombinations("23"))
}

func letterCombinations(A string)  []string {

	chMap := make(map[byte][]byte)

	chMap['0']=[]byte{'0'}
	chMap['1']=[]byte{'1'}
	chMap['2']=[]byte{'a','b','c'}
	chMap['3']=[]byte{'d','e','f'}
	chMap['4']=[]byte{'g','h','i'}
	chMap['5']=[]byte{'j','k','l'}
	chMap['6']=[]byte{'m','n','o'}
	chMap['7']=[]byte{'p','q','r','s'}
	chMap['8']=[]byte{'t','u','v'}
	chMap['9']=[]byte{'w','x','y','z'}

	result := [][]byte{[]byte{}}
	for i:=0;i<len(A);i++ {
		temp := [][]byte{}
		for _, suff := range result {
			for _, b := range chMap[A[i]] {
				suffCopy := make([]byte, len(suff))
				copy(suffCopy, suff)
				suffCopy = append(suffCopy, b)
				temp = append(temp, suffCopy)
			}
		}
		result = temp
	}


	strResult := []string{}

	for _, arr := range result {
		strResult = append(strResult, string(arr))
	}

	return strResult
}

