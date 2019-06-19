package main

import (
	"container/list"
	"fmt"
	"math"
)

func main() {
	slice := []int{7, 9, 5, 11, 8}
	quicksort(slice, 0, len(slice)-1)
	fmt.Println(slice)
	n := 13
	fmt.Println(numSquares(n))
}

func quicksort(slice []int, i int, j int) {
	if i >= j {
		return
	}
	base := slice[i]
	start, end := i, j
	for i < j {
		if base <= slice[j] {
			j--
		} else {
			slice[i], slice[j] = slice[j], slice[i+1]
			i++
		}
	}
	slice[i] = base
	quicksort(slice, start, i)
	quicksort(slice, i+1, end)
}

func numSquares(n int) int {
	base := int(math.Sqrt(float64(n)))
	queue := list.New().Init()
	sum := 0
	res := 0
	for i := base; i >= 1; {
		queue.PushBack(i)
		back := queue.Back().Value.(int)
		if sum+(back*back) > n {
			i--
			continue
		} else {
			sum = sum + (back * back)
			res++
			queue.Remove(queue.Back())
			if sum+back == n {
				return res
			}
		}
	}
	return res
}
