package main

import (
	"fmt"
	"strconv"
)

type Post struct {
	data int
	next *Post
}

func n(node **Post, data int) {
	temp := *node
	w := new(Post)
	w.data = data
	w.next = nil
	if *node == nil {
		*node = w
		return
	}
	for temp.next != nil {

		temp = temp.next
	}
	temp.next = w
}
func print_list(node **Post) {
	temp := *node
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}
func swap(a *Post, s *Post) {
	d := new(Post)
	d.data = a.data
	a.data = s.data
	s.data = d.data
}
func bubbleSort(start **Post) {
	var swapped bool
	ptr1 := new(Post)
	ptr1 = nil
	lptr := new(Post)
	lptr = nil
	if (*start) == nil {
		return
	}
	for true {
		swapped = false
		ptr1 = *start
		for ptr1.next != lptr {
			if ptr1.data > ptr1.next.data {
				the_next := ptr1.next
				swap(ptr1, the_next)
				swapped = true
			}
			ptr1 = ptr1.next
		}
		lptr = ptr1
		if swapped {
			continue
		} else {
			break
		}
	}
}
func main() {
	p1 := new(Post)
	p1 = nil
	for true {
		var x string
		fmt.Scan(&x)
		if x == "x" {
			break
		}
		i2, _ := strconv.Atoi(x)

		n(&p1, i2)
	}
	bubbleSort(&p1)
	print_list(&p1)
}
