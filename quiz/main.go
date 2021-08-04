package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	//This flag is if you use the -h or --help, not exactly needed for this, but good to know if you're using pure binary files etc.
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	//If it can't open the file (As in it doesn't exist), informs the user of the error then exits out the application.
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	//if it can't read/parse the CSV even though it exists, this error is thrown and the program exits.
	if err != nil {
		exit("Failed to parse provided CSV file.")
	}
	fmt.Println(lines)
}

//this function will read in our CSV and highlight problems up to the CSV length
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	// Not using append here. We know the size of what we're parsing so we don't need to accomodate for if it's larger/shorter.
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

type problem struct {
	// here we're telling the program what to expect when it sees a problemtype
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
