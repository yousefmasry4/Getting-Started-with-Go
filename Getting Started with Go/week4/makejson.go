package main

import "encoding/json"
import "fmt"

func main() {
	m := make(map[string]string)
	var x, c string
	fmt.Printf("enter the name:")
	fmt.Scanln(&x)
	fmt.Print("enter the address:")
	fmt.Scanln(&c)
	fmt.Print("json:")
	m["name"] = x
	m["address"] = c
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	jsonStr := string(data)
	fmt.Println(jsonStr)
	fmt.Println(m)

}
