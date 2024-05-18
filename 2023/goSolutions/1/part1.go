package main

import (
	"fmt"
	"os"
	"strconv"
)

func Part1() {
	// read file contents
	input, err := os.ReadFile(inputFilename)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}

	// init variables
	sum, temp := 0, 0

	// main loop
	for _, char := range input {
		if string(char) == "\n" {
			sum += temp
			temp = 0
		}

		num, err := strconv.Atoi(string(char))
		if err == nil {
			if temp == 0 {
				sum += num * 10
			}
			temp = num
		}
	}

	fmt.Println(sum)

}
