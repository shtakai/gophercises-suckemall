package quiz

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func LoadCsv() {
	quizes := []string {
		"1+1, 2",
		"2+1, 3",
	}
	fmt.Printf("load file")

	csvString := strings.Join(quizes, "\n")

    r := csv.NewReader(strings.NewReader(csvString))
    for {
        record, err := r.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Println("Read error: ", err)
            break
        }

        fmt.Printf("%s = %s\n", record[0], record[1])
    }
}