package main

import (
	day1 "adventofcode/v2015"
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
			day1.Run()
		default:
			log.Panicln("Day not implemented")
		}
	default:
		log.Panicln("Year not implemented")
	}
}
