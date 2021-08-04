package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	//This flag is if you use the -h or --help, not exactly needed for this, but good to know if you're using pure binary files etc.

	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
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
	//we can print the lines, our quiz goes from a 2d slice to just slices structured with our values.

	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0

	//this will print every problem, without giving the answers. Also proceeds them with Problem #d, spicy stuff.

	for i, p := range problems {
		select {
		case <-timer.C:
			fmt.Printf("Time limit reached. You got %d correct out of %d.\n", correct, len(problems))
			return
		default:
			fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
			var answer string

			//this gets rid of all spaces, works for numerical and single word quizzes, not for ones that require multi-line answers. We've also given a pointer to the answer variable so that it knows what to expect.

			fmt.Scanf("%s\n", &answer)

			if answer == p.a {
				correct++
			} else {
				fmt.Printf("WRONG! The correct answer is %s.\n", p.a)
			}
		}

	}
	fmt.Printf("You got %d correct out of %d.\n", correct, len(problems))
}

//this function will read in our CSV and highlight problems up to the CSV length
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	// Not using append here. We know the size of what we're parsing so we don't need to accomodate for if it's larger/shorter.

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],

			//dirty fix to trimspace and help validate the CSV.

			a: strings.TrimSpace(line[1]),
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
