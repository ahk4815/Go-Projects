package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to Quiz App")

	fileName := os.Args[1]

	fmt.Println(fileName)
}
