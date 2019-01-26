package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
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

	reader := csv.NewReader(fp)
	reader.LazyQuotes = true
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Println(record)
	}

}
