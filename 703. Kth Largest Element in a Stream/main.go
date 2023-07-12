package main

import "container/heap"

type KthLargest struct {
	minHeap *MinHeap
	K       int
}

func Constructor(k int, nums []int) KthLargest {
	kth := KthLargest{minHeap: &MinHeap{}, K: k}

	for _, num := range nums {
		kth.Add(num)
	}

	return kth
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.minHeap, val)

	for this.minHeap.Len() > this.K {
		heap.Pop(this.minHeap)
	}

	return (*this.minHeap)[0]
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

func main() {

}
