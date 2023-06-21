package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type NodeL struct {
	Head *ListNode
	Tail *ListNode
}

func pushBack(l *NodeL, val int) {
	newN := &ListNode{Val: val}
	if l.Head == nil && l.Tail == nil {
		l.Head = newN
		l.Tail = newN
	} else {
		l.Tail.Next = newN
		l.Tail = newN
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	node := head
	next := node.Next
	for {
		if node.Next != nil {
			if node.Val == next.Val {
				node.Next = next.Next
				next = node.Next
				continue
			}
			node = next
		} else {
			break
		}
	}
	return head
}

func main() {
	numbers := []int{1, 1, 2, 3, 3}
	l1 := &NodeL{}
	for _, val := range numbers {
		pushBack(l1, val)
	}
	l2 := deleteDuplicates(l1.Head)
	for i := l2; i != nil; i = i.Next {
		fmt.Println(i.Val)
	}
}
