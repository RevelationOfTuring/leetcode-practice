package main

import (
	"fmt"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	//input, s := []int{2, 3, 1, 2, 4, 3}, 7
	input, s := []int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8}, 15
	fmt.Println(minSubArrayLen(s, input))
}
