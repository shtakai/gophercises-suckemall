package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type QAndA struct {
	Question string
	Answer   int
}

var correct = 0

func main() {
	numberOfQuestions, filename, timer := setup()

	var fp *os.File
	fp, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	var qAndAs = make([]QAndA, 0)
	qAndAs = createQAndAs(fp, qAndAs)

	go performQAndA(numberOfQuestions, qAndAs)

	time.Sleep(time.Duration(*timer) * time.Second)
	fmt.Printf("Your fucked result is %v / %v\n", correct, *numberOfQuestions)
}

func setup() (*int, *string, *int) {
	var numberOfQuestions = flag.Int("q", 100, "number of questions")
	var filename = flag.String("f", "problems.csv", "question file")
	var timer = flag.Int("t", 30, "timer")
	flag.Parse()
	return numberOfQuestions, filename, timer
}

func performQAndA(numberOfQuestions *int, qAndAs []QAndA) {
	for i := 0; i < *numberOfQuestions; i++ {
		stdInReader := bufio.NewReader(os.Stdin)
		randomIndex := rand.Intn(len(qAndAs))
		fmt.Printf("Question: %v = ?:", qAndAs[randomIndex].Question)
		text, _ := stdInReader.ReadString('\n')
		answer, _ := strconv.Atoi(strings.TrimSpace(text))
		if answer == qAndAs[randomIndex].Answer {
			fmt.Println("correct\n")
			correct++
		} else {
			fmt.Println("fuck incorrect \n")
		}
	}
}

func createQAndAs(fp *os.File, qAndAs []QAndA) []QAndA {
	reader := csv.NewReader(fp)
	reader.LazyQuotes = true
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		answer, err := strconv.Atoi(record[1])
		if err != nil {
			panic(err)
		}
		qAndAs = append(qAndAs, QAndA{Question: record[0], Answer: answer})
	}
	return qAndAs
}
