package main

import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

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

func findKthLargest(nums []int, k int) int {
	h := &MinHeap{}
	for i := 0; i < len(nums); i++ {
		if h.Len() < k {
			heap.Push(h, nums[i])
		} else if (*h)[0] < nums[i] {
			heap.Pop(h)
			heap.Push(h, nums[i])
		}
	}
	return (*h)[0]
}

func main() {

}
