package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

/// Runs the solution for day one
func RunDayOne(in string) {
	var f = openFile(in)
	var elves = totalCalPerElf(f)
	// Take the 3 elves with the highest calories, find their total cals
	sort.Slice(elves, func(i int, j int) bool {
		return elves[i] > elves[j]
	})
	var total = elves[0] + elves[1] + elves[2]

	fmt.Println("-- Day One --")
	fmt.Println("(Part One) Highest Cals:", elves[0])
	fmt.Println("(Part Two) Sum of Highest 3:", total)
}

func openFile(in string) *os.File {
	f, err := os.Open(in)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	return f
}

func totalCalPerElf(f *os.File) []int {
	scanner := bufio.NewScanner(f)
	var elves []int
	currElf := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			elves = append(elves, currElf)
			currElf = 0
		}

		val, _ := strconv.Atoi(scanner.Text())
		currElf += val
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return elves
}
