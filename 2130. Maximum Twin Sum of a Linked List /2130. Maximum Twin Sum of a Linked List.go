package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	slow := head
	var prev *ListNode

	for fast := head; fast != nil && fast.Next != nil; {
		fast = fast.Next.Next

		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	return maxSum(prev, slow)
}

func maxSum(n1, n2 *ListNode) int {
	res := 0
	for i, j := n1, n2; i != nil && j != nil; {
		tmp := i.Val + j.Val
		if res < tmp {
			res = tmp
		}
		i = i.Next
		j = j.Next
	}

	return res
}

func main() {
	head := &ListNode{}
	count := 0
	for i := head; count < 10; i = i.Next {
		count++
		i.Val = count
		i.Next = &ListNode{}
		if count == 10 {
			i.Next = nil
		}
	}

	fmt.Println(pairSum(head))
}
