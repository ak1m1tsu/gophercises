package main

import "github.com/romankravchuk/learn-go/quiz"

func main() {
	config := quiz.NewConfig()
	file := quiz.GetCsvFile(*config.CsvFilename)
	lines := quiz.ReadCsvFile(file)
	problems := quiz.ParseLines(lines, *config.IsShuffle)
	timer := quiz.GetTimer(config.TimeLimit)
	quiz.PrintCountCorrectAnswers(problems, timer)
}
