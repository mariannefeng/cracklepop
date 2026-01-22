package main

import (
	"fmt"
	"os"
	"strconv"
)

// Prints numbers 1 to MAX_NUM (default 100).
// Multiples of 3 print "Crackle", multiples of 5 print "Pop", multiples of both print "CracklePop".
func main() {
	maxNum, _ := strconv.Atoi(os.Getenv("MAX_NUM"))
	if maxNum <= 0 {
		maxNum = 100
	}

	for i := 1; i <= maxNum; i++ {
		output := ""

		if i%3 == 0 {
			output += "Crackle"
		}

		if i%5 == 0 {
			output += "Pop"
		}

		if output == "" {
			output = strconv.Itoa(i)
		}

		fmt.Println(output)
	}
}
