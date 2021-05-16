package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world!")
	
	res, err := Add(4, 5)
	fmt.Println(res, err)
}