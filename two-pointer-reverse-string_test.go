package main

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	//input:=[]byte("hello")
	input:=[]byte("hannah")
	reverseString(input)
	fmt.Println(input)
}
