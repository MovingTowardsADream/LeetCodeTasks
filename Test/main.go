package main

import (
	"context"
	"fmt"
	"time"
)

func gen(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, val := range nums {
			select {
			case out <- val:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func sq(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			time.Sleep(170 * time.Millisecond)
			select {
			case <-ctx.Done():
				return
			case val, ok := <-in:
				if !ok {
					return
				}
				select {
				case out <- val * val:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return out
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	for val := range sq(ctx, gen(ctx, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)) {
		fmt.Println(val)
	}
}

//
//type Node struct {
//	Key, Value int
//	Prev, Next *Node
//}
//
//type LRUCache struct {
//	Hash          map[int]*Node
//	First, Last   *Node
//	Len, Capacity int
//}
//
//func Constructor(capacity int) LRUCache {
//	return LRUCache{
//		Hash:     map[int]*Node{},
//		First:    nil,
//		Last:     nil,
//		Len:      0,
//		Capacity: capacity,
//	}
//}
//
//func (this *LRUCache) Get(key int) int {
//	val, ok := this.Hash[key]
//	if !ok {
//		return -1
//	}
//	tmp_left, tmp_right := val.Prev, val.Next
//	if tmp_right == nil {
//		return val.Value
//	} else if tmp_left == nil {
//		this.First = this.First.Next
//		this.Last.Next = val
//		this.Last = val
//	} else {
//		tmp_left.Next = tmp_right
//		tmp_right.Prev = tmp_left
//		this.Last.Next = val
//		this.Last = val
//	}
//	return val.Value
//}
//
//func (this *LRUCache) Put(key int, value int) {
//	_, ok := this.Hash[key]
//	if !ok {
//		if this.Len == 0 {
//			tmp := &Node{Key: key, Value: value, Prev: nil, Next: nil}
//			this.First, this.Last = tmp, tmp
//			this.Hash[key] = tmp
//		} else if this.Len == this.Capacity {
//			delete(this.Hash, this.First.Key)
//			this.First = this.First.Next
//			if this.First != nil {
//				this.First.Prev = nil
//			}
//			this.Len -= 1
//			tmp := &Node{Key: key, Value: value, Prev: this.Last, Next: nil}
//			this.Last.Next = tmp
//			this.Last = tmp
//			this.Hash[key] = tmp
//		} else {
//			tmp := &Node{Key: key, Value: value, Prev: this.Last, Next: nil}
//			this.Last.Next = tmp
//			this.Last = this.Last.Next
//			this.Hash[key] = tmp
//		}
//		this.Len++
//	} else {
//
//	}
//}
//
//func (this *LRUCache) Return() {
//	fmt.Println(this.Hash)
//	fmt.Println(this.Len)
//	fmt.Println(this.Capacity)
//	fmt.Println(this.First)
//	fmt.Println(this.Last)
//}
//
//func main() {
//	lRUCache := Constructor(3)
//	lRUCache.Put(1, 1)
//	lRUCache.Put(2, 2)
//	lRUCache.Put(3, 3)
//	lRUCache.Put(4, 4)
//	fmt.Println(lRUCache.Get(4))
//	fmt.Println(lRUCache.Get(3))
//	fmt.Println(lRUCache.Get(2))
//	fmt.Println(lRUCache.Get(1))
//	lRUCache.Put(5, 5)
//	fmt.Println(lRUCache.Get(1))
//	fmt.Println(lRUCache.Get(2))
//	fmt.Println(lRUCache.Get(3))
//	fmt.Println(lRUCache.Get(4))
//	fmt.Println(lRUCache.Get(5))
//	lRUCache.Return()
//}
