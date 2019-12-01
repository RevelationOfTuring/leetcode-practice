package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestStrStr(t *testing.T) {
	//input1,input2:="hello","a"
	input1, input2 := "hello", "ll"
	fmt.Println(strings.Index(input1, input2))
}
