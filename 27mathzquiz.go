package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type mathsquestion struct {
	questiontext string
	answer       string
}

func main() {
	csvfilename := flag.String("csv", "questions.csv", "a csv file in format of maths_question,answer(without space)")
	flag.Parse()
	_ = csvfilename
	//name of file which you wish to read from cla
	fmt.Println(*csvfilename)

	//open the file whose name is *csvfilename
	file, error := os.Open(*csvfilename)
	if error != nil {
		fmt.Println("Failed to open file: ", *csvfilename)
		os.Exit(1)
	}

	//read the file
	pointerReader := csv.NewReader(file)
	//rad all lines from pointer reference to file
	data, errorReader := pointerReader.ReadAll()
	if errorReader != nil {
		fmt.Println("Failed to parse CSV file")
		os.Exit(1)
	}
	//if read was success
	//data is 2d array
	fmt.Println(data)
	//convert each value in 2d array to struct instance
	allquestions := convert2darraytostructinstance(data)
	//print the allquestion and see that it is array of type struct instance
	fmt.Println(allquestions)

	//loop on allquestion and ask user for the valid answer
	correctanswercount := 0
	for index, questructinstance := range allquestions {
		fmt.Printf("Question %v. %v\n", index+1, questructinstance.questiontext)
		var useranswer string
		fmt.Scanf("%s\n", &useranswer)
		if useranswer == questructinstance.answer {
			correctanswercount++
		}
	}
	fmt.Printf("Correct answer is %v out of %v", correctanswercount, len(allquestions))
}

func convert2darraytostructinstance(queandans [][]string) []mathsquestion {

	slicequestion := make([]mathsquestion, len(queandans))
	for index, que := range queandans {
		slicequestion[index] = mathsquestion{
			questiontext: que[0],
			answer:       strings.TrimSpace(que[1]),
		}
	}

	return slicequestion

}
