package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var (
		input []int
		i     string
	)
	input = make([]int, 0, 3)
	for {
		fmt.Println("Enter a Number")
		fmt.Scanf("%s", &i)
		if i == "X" {
			break
		}
		in, _ := strconv.ParseInt(i, 10, 32)
		input = append(input, int(in))
		sort.Ints(input)
		fmt.Println(input)
	}
}
