package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {
	if err := mainNoExit(); err != nil {
		log.Fatalf("error: %+v", err)
	}
}

func mainNoExit() error {
	flagCsvFilename := flag.String("csv", "projects/go_quiz/problems.csv", "a csv file in the format of 'question,answer'")
	flagTimeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flagRandom := flag.Bool("random", false, "randomize order of questions")
	flag.Parse()
	_ = flagRandom

	lines, err := readInputFile(*flagCsvFilename)
	if err != nil {
		return fmt.Errorf("failed while reading input csv file: %w", err)
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*flagTimeLimit) * time.Second)
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			_, _ = fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
			return nil
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))

	return nil
}

func readInputFile(filename string) ([][]string, error) {
	csvPath, err := filepath.Abs(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to parse path: %w", err)
	}
	file, err := os.Open(csvPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open the CSV file: %w", err)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to parse the provided CSV file: %w", err)
	}
	return lines, nil
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: strings.TrimSpace(line[0]),
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}
