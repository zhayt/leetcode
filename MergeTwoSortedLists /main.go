package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkList struct {
	Head *ListNode
	Tail *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	list := ListNode{}
	link := &list
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			link.Next = list1
			list1 = list1.Next
		} else {
			link.Next = list2
			list2 = list2.Next
		}
		link = link.Next
	}
	if list1 == nil {
		link.Next = list2
	}
	if list2 == nil {
		link.Next = list1
	}
	return list.Next
}

func main() {
	link1 := &LinkList{}
	link2 := &LinkList{}
	nums1 := []int{1, 2, 4}
	nums2 := []int{3, 5, 6}
	toRange(link1, nums1, ListPushBack)
	toRange(link2, nums2, ListPushBack)

	printerLinkedList(mergeTwoLists(link1.Head, link2.Head))
}

func printerLinkedList(l *ListNode) {
	for i := l; i != nil; i = i.Next {
		fmt.Println(i.Val)
	}
}

func toRange(l *LinkList, arr []int, f func(list *LinkList, num int)) {
	for _, val := range arr {
		f(l, val)
	}
}

func ListPushBack(l *LinkList, data int) {
	lNode := &ListNode{Val: data}
	if l.Tail == nil {
		l.Head = lNode
		l.Tail = lNode
	} else {
		l.Tail.Next = lNode
		l.Tail = lNode
	}
}
