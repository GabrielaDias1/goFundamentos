package main

import "fmt"

func main() {
	fmt.Println("1")
	defer fmt.Println("2")//ultima
	fmt.Println("3")
}
