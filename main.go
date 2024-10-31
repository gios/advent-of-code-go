package main

import (
	v2015 "adventofcode/v2015"
	"flag"
	"log"
)

func main() {
	year := flag.Int("year", 2015, "Year of the puzzle")
	day := flag.Int("day", 1, "Day of the puzzle")

	flag.Parse()

	switch *year {
	case 2015:
		switch *day {
		case 1:
			v2015.Day1()
		case 2:
			v2015.Day2()
		default:
			log.Panicln("Day not implemented")
		}
	default:
		log.Panicln("Year not implemented")
	}
}
