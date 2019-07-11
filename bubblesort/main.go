package main

import "fmt"

func main() {
	li := []int{2, 5, 1, 7, 4, 11, 8, 3}
	bubblesort(li)
	fmt.Println(li)
}

func bubblesort(li []int) {
	for i := 0; i < len(li); i++ {
		for j := 0; j < len(li)-1-i; j++ {
			if li[j] > li[j+1] {
				li[j+1], li[j] = li[j], li[j+1]
			}
		}
	}
}
