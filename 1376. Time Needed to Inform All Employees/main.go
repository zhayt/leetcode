package main

import (
	"container/list"
	"fmt"
)

type track struct {
	informTime int
	employee   int
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	graph := make(map[int][]int, n)
	for i, m := range manager {
		graph[m] = append(graph[m], i)
	}

	var res int
	q := list.New()
	q.PushBack(&track{
		informTime: 0,
		employee:   headID,
	})
	for q.Len() > 0 {
		elem := q.Front()
		curr := elem.Value.(*track)
		if curr.informTime > res {
			res = curr.informTime
		}
		q.Remove(elem)
		for _, v := range graph[curr.employee] {
			q.PushBack(&track{
				informTime: curr.informTime + informTime[curr.employee],
				employee:   v,
			})
		}
	}

	return res
}

func main() {
	fmt.Println(numOfMinutes(22, 7, []int{12, 7, 18, 11, 13, 21, 12, -1, 6, 5, 14, 13, 14, 9, 20, 13, 11, 1, 1, 2, 3, 19}, []int{0, 540, 347, 586, 0, 748, 824, 486, 0, 777, 0, 202, 653, 713, 454, 0, 0, 0, 574, 735, 721, 772}))
}
