package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 20, "the time limit for the quiz")
	flag.Parse()
	
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("failed to open %s\n", *csvFileName))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("failed to parse the provided csv file: %s", *csvFileName))
	}

	problems := parseLines(lines)
	count := 0
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	problemLoop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i + 1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			break problemLoop
		case answer := <-answerCh:
			if answer == p.a {
				count += 1
			}
		}
	}
	fmt.Printf("\nYou got %d problems correct out of %d", count, len(problems))
}