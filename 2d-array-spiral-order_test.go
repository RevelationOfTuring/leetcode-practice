package main

import (
	"fmt"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	//input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// [1 2 3 6 9 8 7 4 5]
	//input:=[][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}}
	// [1 2 3 4 8 12 11 10 9 5 6 7]
	input:=[][]int{{1,2,3,4}}

	fmt.Println(spiralOrder(input))

}
