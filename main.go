package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	q string
	a string
}

func main() {
	// flag
	// read the csv
	// parse the lines
	// assing to the struct
	// ask question and get the answer
	// print score

	fileName := flag.String("csv", "porblem.csv", "a csv file in the format od questions and numbers")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	problems := ParseLine(lines)

	reader := bufio.NewReader(os.Stdin)
	count := 0
	for i, p := range problems {

		fmt.Printf("PROBLEM #%d. %s = ", i+1, p.q)
		anser, _ := reader.ReadString('\n')

		if strings.TrimSpace(anser) == p.a {
			// fmt.Println("Correct!")
			count++
		}

	}

	fmt.Printf("\nYour score is: %d out of %d\n", count, len(problems))

}

func ParseLine(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{q: line[0], a: line[1]}
	}
	return problems

}
