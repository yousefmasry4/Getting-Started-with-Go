package main

import "fmt"

func main() {
	flag := 1
	yes := 0
	var x string
	fmt.Scan(&x)
	l := len(x)
	for i := 0; i < l-1; i++ {
		if x[i] == ' ' {
			continue
		}
		if flag == 1 {
			if x[i] == 'i' || x[i] == 'I' {
				yes = yes + 1
			}
			flag = 0
		}
		if flag == 0 {
			if x[i] == 'a' || x[i] == 'A' {
				yes = yes + 1
				break
			}
		}
	}
	if x[l-1] == 'n' || x[l-1] == 'N' {
		yes = yes + 1
	}
	if yes == 3 {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("NOT Found!")
	}
}
