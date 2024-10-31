package v2015

import (
	u "adventofcode/utils"
	"fmt"
	"log"
)

func Day1() {
	input, err := u.GetInput(2015, 1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("2015: Day 1, Part 1: Not Quite Lisp")
	fmt.Println(floor(input))
	fmt.Println("2015: Day 1, Part 2: Not Quite Lisp")
	fmt.Println(basement(input))
}

func floor(input string) int {
	var counter int

	for _, r := range input {
		if r == '(' {
			counter++
		} else {
			counter--
		}
	}

	return counter
}

func basement(input string) int {
	var counter int

	for i, r := range input {
		if r == '(' {
			counter++
		} else {
			counter--
		}

		if counter == -1 {
			return i + 1
		}
	}

	return -1
}
