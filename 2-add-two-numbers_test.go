package main

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := createList(1)
	l2 := createList(9,9)
	ret := addTwoNumbers(l1, l2)
	showListNode(ret)

}

func createList(ints ...int) *ListNode {
	var p = &ListNode{}
	var head = p
	l := len(ints)
	for i := 0; i < l; i++ {
		p.Val = ints[i]
		if i != l-1 {
			p.Next = &ListNode{}
			p = p.Next
		} else {
			break
		}
	}
	return head
}

func showListNode(head *ListNode) {
	for ; head != nil; head = head.Next {
		fmt.Println(head.Val)
	}
}
