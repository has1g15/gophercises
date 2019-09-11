package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
)

func main() {
	programArguments := os.Args[1:]

	for _, element := range programArguments{
		if element == "" {
			break
		} else if element == "help" {
			fmt.Println("Helpful message")
		} else {
			fmt.Printf("You have %s seconds to complete quiz\n", element)
		}
	}

	fmt.Printf("Welcome to the Quiz\n")

	csvFile, _ := os.Open("problems.csv")
	csvReader := csv.NewReader(csvFile)
	stdinReader := bufio.NewReader(os.Stdin)

	sumCorrect := 0
	sumIncorrect := 0
	var quiz []*qa

	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("Could not open file", err)
		}
		questionAnswer := new(qa)
		questionAnswer.question = line[0]
		questionAnswer.answer = line[1]
		quiz = append(quiz, questionAnswer)
	}

	for len(quiz) > 0 {
		questionAnswer, index := getRandomQuestion(quiz)

		//Ask question
		fmt.Printf("What is %s?\n", questionAnswer.question)

		//Read in answer
		text, _ := stdinReader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")

		//Compare given answer to actual answer
		if text == questionAnswer.answer {
			fmt.Printf("Correct\n")
			sumCorrect++
		} else {
			fmt.Printf("Incorrect\n")
			sumIncorrect++
		}

		//Remove asked question from slice
		copy(quiz[index:], quiz[index+1:])
		quiz[len(quiz)-1] = nil
		quiz = quiz[:len(quiz)-1]
	}

	fmt.Printf("You got %d correct answer and %d incorrect answer", sumCorrect, sumIncorrect)
}

type qa struct {
	question string
	answer string
}

func getRandomQuestion(quiz[]*qa) (*qa, int){
	random := rand.Intn(len(quiz))
	return quiz[random], random
}