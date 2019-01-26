package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	var numberOfQuestions = flag.Int("q", 3, "number of questions")
	var filename = flag.String("f", "problems.csv", "question file")
	flag.Parse()
	fmt.Println(filename)
	fmt.Println(*numberOfQuestions)

	var fp *os.File
	fp, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	qAndAs := make(map[string]int)

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
		qAndAs[record[0]] = answer
	}
	for i := 0; i < *numberOfQuestions; i++ {
	}

}
