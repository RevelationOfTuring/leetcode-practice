package main

import (
	"fmt"
	"testing"
)

func TestTwoSum1(t *testing.T) {
	numbers, target := []int{2, 7, 11, 15}, 9
	fmt.Println(twoSum1(numbers, target))
}
