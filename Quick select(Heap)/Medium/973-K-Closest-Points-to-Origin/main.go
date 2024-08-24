package main

import (
	"container/heap"
	"math"
)

// v1

type MaxHeap [][]int

func (h MaxHeap) Len() int { return len(h) }

func (h MaxHeap) Less(i, j int) bool {
	return math.Sqrt(float64(h[i][0]*h[i][0]+h[i][1]*h[i][1])) > math.Sqrt(float64(h[j][0]*h[j][0]+h[j][1]*h[j][1]))
}

func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	h := &MaxHeap{}

	for i := 0; i < len(points); i++ {
		if h.Len() < k {
			heap.Push(h, points[i])
		} else {
			tmp := (*h)[0]
			if math.Sqrt(float64(tmp[0]*tmp[0]+tmp[1]*tmp[1])) > math.Sqrt(float64(points[i][0]*points[i][0]+points[i][1]*points[i][1])) {
				heap.Pop(h)
				heap.Push(h, points[i])
			}
		}
	}
	result := make([][]int, 0, k)
	for i := 0; i < k; i++ {
		result = append(result, (heap.Pop(h)).([]int))
	}
	return result
}

// or v2

//type Point struct {
//	X, Y int
//	Result float64
//}
//
//func NewPoint(x, y int) *Point {
//	return &Point{X: x, Y: y, Result: math.Sqrt(float64(x * x + y * y))}
//}
//
//type MaxHeap []*Point
//
//func (h MaxHeap) Len() int { return len(h) }
//
//func (h MaxHeap) Less(i, j int) bool {
//	return h[i].Result > h[j].Result
//}
//
//func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
//
//func (h *MaxHeap) Push(x interface{}) {
//	*h = append(*h, x.(*Point))
//}
//
//func (h *MaxHeap) Pop() interface{} {
//	old := *h
//	n := len(old)
//	x := old[n-1]
//	*h = old[:n-1]
//	return x
//}
//
//func kClosest(points [][]int, k int) [][]int {
//	h := &MaxHeap{}
//
//	for i := 0; i < len(points); i++ {
//		if h.Len() < k {
//			heap.Push(h, NewPoint(points[i][0], points[i][1]))
//		} else {
//			tmp := NewPoint(points[i][0], points[i][1])
//			if (*h)[0].Result > tmp.Result {
//				heap.Pop(h)
//				heap.Push(h, tmp)
//			}
//		}
//	}
//	result := make([][]int, 0, k)
//	for i := 0; i < k; i++ {
//		tmp := (heap.Pop(h)).(*Point)
//		result = append(result, []int{tmp.X, tmp.Y})
//	}
//	return result
//}

func main() {
}
