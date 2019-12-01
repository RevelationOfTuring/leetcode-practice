package main

import (
	"fmt"
	"testing"
)

func TestArrayPairSum(t *testing.T) {
	//input:=[]int{1,5,23,4,8,3,43,21,3}
	//quickSort(0,len(input)-1,input)
	//fmt.Println(input)

	//input := []int{1, 5, 23, 4, 8, 3, 43, 21, 3}
	//numsLen := len(input)
	//for i := 0; i < numsLen-1; i++ {
	//	for j := 0; j < numsLen-i-1; j++ {
	//		if input[j+1] < input[j] {
	//			input[j+1], input[j] = input[j], input[j+1]
	//		}
	//	}
	//}
	//fmt.Println(input)

	input := []int{1, 4, 3, 2}
	fmt.Println(arrayPairSum(input))
}
