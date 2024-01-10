package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"flag"
	"strings"
)

func main() {
	csvFilename := flag.String("file","problems.csv","Specify the problems file")
	flag.Parse()
	
	file,err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open %s file\n",*csvFilename))
	}
	defer file.Close()

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to read the given csv file")
	}

	correct := 0
	problems := parseLines(&lines)
	for i,p := range *problems {
		fmt.Printf("Problem #%d = %s\n",i+1,p.q)
		var answer string
		fmt.Scanf("%s\n",&answer)
		if p.a == answer {
			correct++
		}
	}

	fmt.Printf("You got %d out of %d",correct,len(*problems))
	
}

type problem struct {
	q string
	a string
}

func parseLines(lines *[][]string) *[]problem {
	ret := make([]problem, len(*lines))
	for i,line := range *lines {
		ret[i] = problem{
			q : line[0],
			a : strings.TrimSpace(line[1]),
		}
	}
	return &ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}