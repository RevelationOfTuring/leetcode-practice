package main

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	//input:=[]string{"flower","flow","flight"}
	input:=[]string{"dog","racecar","car"}
	fmt.Println(longestCommonPrefix(input))
}
