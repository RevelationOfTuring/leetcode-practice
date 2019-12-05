package main

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	//input, k := []int{1, 2, 3, 4, 5, 6, 7}, 3
	input, k := []int{-1, -100, 3, 99}, 2
	//reverseNums(input)
	//fmt.Println(input)


	rotate(input, k)
	fmt.Println(input)




}
