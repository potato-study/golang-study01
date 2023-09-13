package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: lesson04 <name>")
		os.Exit(1)
	}

	size, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Wrong size")
		os.Exit(1)
	}

	for i := 0; i < size; i++ {
		go printAny("Hello,", i)
	}

	time.Sleep(1 * time.Second)
}

func printAny(any ...interface{}) {
	fmt.Println(any...)
}
