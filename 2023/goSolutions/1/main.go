package main

import (
	"os"
)

const inputFilename = "input.txt"

func main() {
	if os.Args[1] == "1" {
		Part1()
	} else {
		Part2()
	}
}
