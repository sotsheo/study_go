package main

import "fmt"

func main() {
	slice := map[int]string{
		1: "Hoang",
		2: "thanh",
	}
	fmt.Println(slice)
	delete(slice, 2)
	fmt.Println(slice)
}
