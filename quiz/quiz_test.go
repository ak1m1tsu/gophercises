package quiz

import (
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {
	config := NewConfig()

	if *config.CsvFilename != "problems.csv" {
		t.Errorf("Expected csv filename 'problems.csv' by default, but got %s", *config.CsvFilename)
	}

	if *config.TimeLimit != 30 {
		t.Errorf("Expected time limit '30', but got %v", *config.TimeLimit)
	}

	if *config.IsShuffle != false {
		t.Errorf("Expected shuffle flag 'false', but got %v", *config.IsShuffle)
	}
}

func TestGetCsvFile(t *testing.T) {
	const file = "_quiztesting.csv"
	os.Remove(file)
	os.WriteFile(file, []byte{}, 0666)
	csvFile := GetCsvFile(file)
	if csvFile.Name() != file {
		t.Errorf("Expected file '_quiztesting.csv', but got %s", csvFile.Name())
	}
	csvFile.Close()
	os.Remove(file)
}

func TestGetAndReadCsvFile(t *testing.T) {
	const file = "_quiztesting.csv"

	os.Remove(file)
	os.WriteFile(file, []byte("1+1,2"), 0666)

	csvFile := GetCsvFile(file)
	if csvFile.Name() != file {
		t.Errorf("Expected file '_quiztesting.csv', but got %s", csvFile.Name())
	}

	lines := ReadCsvFile(csvFile)
	if lines[0][0] != "1+1" {
		t.Errorf("Expected question '1+1', but got %s", lines[0][0])
	}
	if lines[0][1] != "2" {
		t.Errorf("Expected answer '2', but got %s", lines[0][1])
	}

	csvFile.Close()
	os.Remove(file)
}
