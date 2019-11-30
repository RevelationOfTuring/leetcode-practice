package main

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//链接：https://leetcode-cn.com/problems/add-two-numbers

// example:
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carryBit, bitSum int
	var head = &ListNode{}
	var preP *ListNode
	var retHead = head
	for ; ; l1, l2 = l1.Next, l2.Next {
		if l1 == nil {
			if l2 == nil {
				if carryBit == 1 {
					head.Val = carryBit
					break
				}
				preP.Next = nil
				break
			}
			preP.Next = l2
			head = l2
			for ; carryBit == 1; preP, head = head, head.Next {
				bitSum = head.Val + carryBit
				head.Val = bitSum % 10
				carryBit = bitSum / 10
				if head.Next == nil && carryBit == 1 {
					head.Next = &ListNode{1, nil}
					break
				}
			}

			break

		} else if l2 == nil {
			preP.Next = l1
			head = l1
			for ; carryBit == 1; preP, head = head, head.Next {
				bitSum = head.Val + carryBit
				head.Val = bitSum % 10
				carryBit = bitSum / 10
				if head.Next == nil && carryBit == 1 {
					head.Next = &ListNode{1, nil}
					break
				}
			}

			break
		} else {
			bitSum = l1.Val + l2.Val + carryBit
			head.Val = bitSum % 10
			carryBit = bitSum / 10
			head.Next = &ListNode{}
			preP = head
			head = head.Next
		}
	}
	return retHead
}
