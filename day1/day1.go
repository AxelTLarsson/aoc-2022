package day1

import (
	"fmt"
	"os"
	"sort"
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

	// create a "stack" to store top 3 elve's calories
	var top []int = []int{}

	// elf by elf sum up calories
	for _, elf := range elves {
		calories := strings.Fields(elf)
		// convert to ints and sum up
		sum := 0
		for _, c := range calories {
			calory, err := strconv.Atoi(c)
			check(err)
			sum += calory
		}

		top = append(top, sum)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(top)))
	fmt.Println("top3:", top[:3])
	sum := top[0] + top[1] + top[2]
	fmt.Println("sum:", sum)

}
