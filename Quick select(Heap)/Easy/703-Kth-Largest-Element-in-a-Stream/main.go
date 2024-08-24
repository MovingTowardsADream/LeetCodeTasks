package main

import (
	"container/heap"
)

type MinHeap []int

func (h MinHeap) Len() int { return len(h) }

// Изменение на ">" для MaxHeap
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type KthLargest struct {
	heap *MinHeap
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	h := &MinHeap{}
	heap.Init(h)
	kl := KthLargest{heap: h, k: k}
	for _, value := range nums {
		kl.Add(value)
	}
	return kl
}

func (this *KthLargest) Add(val int) int {
	if this.heap.Len() < this.k {
		heap.Push(this.heap, val)
	} else if val > (*this.heap)[0] {
		heap.Pop(this.heap)
		heap.Push(this.heap, val)
	}

	if this.heap.Len() > 0 {
		return (*this.heap)[0]
	}
	return -1
}

func main() {
}
