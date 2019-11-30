package main

import (
	"fmt"
	"testing"
)

func TestFindDiagonalOrder(t *testing.T) {
	//input:=[][]int{{1, 2, 3},{4, 5, 6},{7, 8, 9}}
	// [1,2,4,7,5,3,6,8,9]
	//input:=[][]int{}
	//input:=[][]int{{1}}
	//input:=[][]int{{2, 3}}
	input := [][]int{{2, 5}, {8, 4}, {0, -1}}

	fmt.Println(findDiagonalOrder(input))


	//2 5
	//8 4
	//0 -1
}
