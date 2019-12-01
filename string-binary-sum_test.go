package main

import (
	"fmt"
	"testing"
)

func TestAddBinary(t *testing.T) {
	var carryBit bool
	fmt.Printf("%c %v\n", bitAdd('0', '0', &carryBit), carryBit)
	carryBit = false
	fmt.Printf("%c %v\n", bitAdd('0', '1', &carryBit), carryBit)
	carryBit = false
	fmt.Printf("%c %v\n", bitAdd('1', '0', &carryBit), carryBit)
	carryBit = false
	fmt.Printf("%c %v\n", bitAdd('1', '1', &carryBit), carryBit)

	carryBit = true
	fmt.Printf("%c %v\n", bitAdd('0', '0', &carryBit), carryBit)
	carryBit = true
	fmt.Printf("%c %v\n", bitAdd('1', '0', &carryBit), carryBit)
	carryBit = true
	fmt.Printf("%c %v\n", bitAdd('0', '1', &carryBit), carryBit)
	carryBit = true
	fmt.Printf("%c %v\n", bitAdd('1', '1', &carryBit), carryBit)
	//testStandardBytes("1111","0")
	testStandardBytes("11", "1")

	input1, input2 := "1010", "1011"
	//input1, input2 := "11", "1"
	fmt.Println(addBinary(input1, input2))
}

func testStandardBytes(a, b string) {
	bytesA, bytesB := []byte(a), []byte(b)
	lenA, lenB := len(bytesA), len(bytesB)
	var standardBytesA, standardBytesB []byte
	if lenA == lenB {
		// 两个数组等长
		standardBytesA, standardBytesB = bytesA, bytesB
	} else {
		// 两个数组不等长
		if lenA < lenB {
			standardBytesB = bytesB
			standardBytesA = make([]byte, lenB)
			for i, j := lenA-1, lenB-1; j >= 0; j -- {
				if i >= 0 {
					standardBytesA[j] = bytesA[i]
					i--
				} else {
					standardBytesA[j] = '0'
				}
			}
		} else {
			standardBytesA = bytesA
			standardBytesB = make([]byte, lenA)
			for i, j := lenB-1, lenA-1; j >= 0; j -- {
				if i >= 0 {
					standardBytesB[j] = bytesB[i]
					i--
				} else {
					standardBytesB[j] = '0'
				}
			}
		}
	}
	fmt.Println(standardBytesA)
	fmt.Println(standardBytesB)
	fmt.Println(len(standardBytesA))
	fmt.Println(len(standardBytesB))
}
