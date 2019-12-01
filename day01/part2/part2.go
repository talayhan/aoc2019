package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var num int
	var subnum int
	var sum int

	for {

		_, err := fmt.Fscanf(file, "%d\n", &num) // give a patter to scan

		if err != nil {

			if err == io.EOF {
				break // stop reading the file
			}
			fmt.Println(err)
			os.Exit(1)
		}

		for subnum = (num / 3) - 2; subnum > 0; subnum = (subnum / 3) - 2 {
			sum += subnum
		}
	}

	// print out the sum
	fmt.Println(sum)
}
