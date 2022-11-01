package quiz

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

type config struct {
	csvFilename *string
	timeLimit   *int
	isShuffle   *bool
}

func Run() {
	config := newConfig()
	file := openAndReadCsvFile(*config.csvFilename)
	lines := readCsvFile(file)
	problems := parseLines(lines, *config.isShuffle)
	timer := getTimer(config.timeLimit)
	printCountCorrectAnswers(problems, timer)
}

func newConfig() config {
	csvFilename := flag.String("csv", "problems.csv", "A csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "The time limit for the quiz in seconds")
	isShuffle := flag.Bool("shuffle", false, "Shuffle problems or not")
	flag.Parse()
	return config{
		csvFilename: csvFilename,
		timeLimit:   timeLimit,
		isShuffle:   isShuffle,
	}
}

func getTimer(limit *int) *time.Timer {
	return time.NewTimer(time.Second * time.Duration(*limit))
}

func openAndReadCsvFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", filename))
	}
	return file
}

func readCsvFile(file *os.File) [][]string {
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	return lines
}

func parseLines(lines [][]string, isShuffle bool) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(strings.ToLower(line[1])),
		}
	}
	if isShuffle {
		shuffleProblems(problems)
	}
	return problems
}

func shuffleProblems(problems []problem) {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range problems {
		newPosition := r.Intn(len(problems) - 1)
		problems[i], problems[newPosition] = problems[newPosition], problems[i]
	}
}

func printCountCorrectAnswers(problems []problem, timer *time.Timer) {
	correct := 0
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.question)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == problem.answer {
				correct++
			}
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
