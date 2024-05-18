package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	// read file contents
	input, err := os.ReadFile(inputFilename)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}

	// init variables
	sum, temp := 0, 0

	digits := make([]string, 9)
	digits[0] = "one"
	digits[1] = "two"
	digits[2] = "three"
	digits[3] = "four"
	digits[4] = "five"
	digits[5] = "six"
	digits[6] = "seven"
	digits[7] = "eight"
	digits[8] = "nine"

	// main loop
	for idx, char := range input {
		if string(char) == "\n" {
			sum += temp
			temp = 0
		}

		for jdx, digit := range digits {
			// caveat: the index should be 0 since we are searching in string[idx:]
			if strings.Index(string(input[idx:]), digit) == 0 {
				if temp == 0 {
					sum += (jdx + 1) * 10
				}
				temp = jdx + 1
			}
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
