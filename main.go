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
	//read values from flags in the command line
	csvFilename := flag.String("file", "problems.csv", "Specify the problems file")
	limit := flag.Int("limit", 30, "Time limit for the quiz")

	//MANDATORY: parse the flags to capture the values
	flag.Parse()

	//Open the file
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open %s file\n", *csvFilename))
	}
	//close the file before the main function exists regardless of any error or success
	defer file.Close()

	//io Reader to read the csv files
	r := csv.NewReader(file)
	//parse the file and read the lines[][]
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to read the given csv file")
	}

	correct := 0
	//get the problems
	problems := parseLines(&lines)


	//Initialize the timer
	timer := time.NewTimer(time.Duration(*limit) * time.Second)

	for i, p := range *problems {
		fmt.Printf("Problem #%d = %s\n", i+1, p.q)
		answerCh := make(chan string)
		//goroutine so as to receive the answer in non blocking manner
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
	
			answerCh <- answer
		}()

		//either time runs out or we get the answer back, if timeout occurs before return main function
		select {
		case <-timer.C:
			fmt.Printf("\nYou got %d out of %d", correct, len(*problems))
			return
		case answer := <-answerCh:
			if p.a == answer {
				correct++
			}
		}
	}

	fmt.Printf("You got %d out of %d", correct, len(*problems))

}

//problem format
type problem struct {
	q string
	a string
}


//function to converthe [][]lines to []problem
func parseLines(lines *[][]string) *[]problem {
	ret := make([]problem, len(*lines))
	for i, line := range *lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return &ret
}

//custom exit function
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
