package v2015

import (
	u "adventofcode/utils"
	"fmt"
	"log"
)

func Run() {
	input, err := u.GetInput(2015, 1)

	if err != nil {
		log.Fatal(err)
	}

	var counter int

	for _, r := range input {
		if r == '(' {
			counter++
		} else {
			counter--
		}
	}

	fmt.Println("Part 1: ", counter)
}
