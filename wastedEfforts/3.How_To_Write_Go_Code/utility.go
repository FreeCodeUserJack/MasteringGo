package main

import (
	"fmt"
	"strconv"
)

func Add(a, b int) (int, int) {
	fmt.Println("Adding " + strconv.Itoa(a) + " and " + strconv.Itoa(b))
	return a + b, 0
}