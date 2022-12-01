package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Run() {
	fmt.Println("Day 1")

	// read data
	dat, err := os.ReadFile("data/day1.txt")
	check(err)

	// split by "elf"
	elves := strings.Split(string(dat), "\n\n")

	// elf by elf sum up calories
	max := 0
	for _, elf := range elves {
		calories := strings.Fields(elf)
		// convert to ints and sum up
		sum := 0
		for _, c := range calories {
			calory, err := strconv.Atoi(c)
			check(err)
			sum += calory
		}

		if sum > max {
			max = sum
		}
	}
	fmt.Println(max)
}
