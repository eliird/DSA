package main

import "fmt"

func main() {
    fmt.Println("hello world")

	//array declaration in go

	var a[10] string

	a[0] = "My"
	a[1] = "Name"

	fmt.Println(a)
	fmt.Println(a[0], a[1], a[2])
	fmt.Println(len(a[0]))
	
	
}