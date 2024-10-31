package v2015

import (
	u "adventofcode/utils"
	"fmt"
	"log"
	"strings"
)

func Day2() {
	input, err := u.GetInput(2015, 2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("2015: Day 2, Part 1: I Was Told There Would Be No Math")
	fmt.Println(wrappingPaper(input))
}

func wrappingPaper(input string) uint64 {
	var total uint64

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		var l, w, h uint64

		fmt.Sscanf(line, "%dx%dx%d", &l, &w, &h)

		total += 2*l*w + 2*w*h + 2*h*l + min(l*w, w*h, h*l)
	}

	return total
}
