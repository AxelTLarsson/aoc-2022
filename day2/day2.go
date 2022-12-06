package day2

import (
	"fmt"
	"os"
	"strings"
)

// outcome
const (
	Loss int = 0
	Draw int = 3
	Win  int = 6
)

type Hand string

const (
	Undefined Hand = ""
	Rock           = "X"
	Paper          = "Y"
	Scissors       = "Z"
	A              = "A"
	B              = "B"
	C              = "C"
)

func (h Hand) String() string {
	switch h {
	case Rock:
		return "Rock"
	case Paper:
		return "Paper"
	case Scissors:
		return "Scissors"
	case A:
		return "A (Rock)"
	case B:
		return "B (Paper)"
	case C:
		return "C (Scissors)"
	}

	return "unknown hand"
}

/*
Paper 2
Rock 1
Scissors 3
*/
func score(them, us Hand) int {
	switch them {
	// Rock
	case A:
		switch us {
		case Rock:
			return Draw + 1
		case Paper:
			return Win + 2
		case Scissors:
			return Loss + 3
		}
	// Paper
	case B:
		switch us {
		case Rock:
			return Loss + 1
		case Paper:
			return Draw + 2
		case Scissors:
			return Win + 3
		}
	// Scissors
	case C:
		switch us {
		case Rock:
			return Win + 1
		case Paper:
			return Loss + 2
		case Scissors:
			return Draw + 3
		}
	}
	return 99999999
}

type strategy struct {
	hand Hand
}

func Run() {
	fmt.Println("Day 2 🎄")

	dat, err := os.ReadFile("data/day2.txt")
	if err != nil {
		panic(err)
	}
	totalScore := 0
	for _, line := range strings.Split(string(dat), "\n") {
		fields := strings.Fields(line)
		if len(fields) == 2 {
			var handOp Hand = Hand(fields[0])
			var handUs Hand = Hand(fields[1])
			outcome := score(handOp, handUs)
			totalScore += outcome
		}
	}
	fmt.Println("total score:", totalScore)
}
