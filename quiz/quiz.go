package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	filename := "problems.csv"
	var numberOfQuestions = flag.Int("q", 3, "number of questions")
	flag.Parse()
	fmt.Println(filename)
	fmt.Println(*numberOfQuestions)

	var fp *os.File
	fp, err := os.Open(filename)
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

	//quizes := []string{
	//	"1+1, 2",
	//	"2+1, 3",
	//}
	//fmt.Printf("load file\n")
	//
	//csvString := strings.Join(quizes, "\n")
	//
	//r := csv.NewReader(strings.NewReader(csvString))
	//for {
	//	record, err := r.Read()
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		fmt.Println("Read error: ", err)
	//		break
	//	}
	//
	//	fmt.Printf("%s = %s\n", record[0], record[1])
	//}
}
