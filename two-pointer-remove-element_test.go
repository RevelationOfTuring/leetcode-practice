package main

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	//input, target := []int{3, 2, 2, 3}, 3
	//input, target := []int{0, 1, 2, 2, 3, 0, 4, 2}, 2
	input, target := []int{ 4, 5}, 4
	fmt.Println(removeElement(input, target))
}
