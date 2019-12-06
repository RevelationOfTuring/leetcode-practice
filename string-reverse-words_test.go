package main
import (
	"fmt"
	"testing"
)

func TestReverseWords(t *testing.T) {
	//input:="the sky is blue"
	input:="a good   example"
	fmt.Println(reverseWords(input))
}
