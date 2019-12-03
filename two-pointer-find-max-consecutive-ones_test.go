package main

import (
	"fmt"
	"testing"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	//input := []int{1, 1, 0, 1, 1, 1}
	input := []int{1,0,1,1,0,1}
	fmt.Println(findMaxConsecutiveOnes(input))
}
