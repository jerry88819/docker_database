package main

import (
	"fmt"
	"testfile/test1"
)


func main() {
	fmt.Println("Test start!")
	var temp int = 1
	fmt.Println(test1.Add(temp))
}