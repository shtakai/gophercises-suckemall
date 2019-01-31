package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/shtakai/gophercises-suckemall/adventure"
)

func main() {
	filename := flag.String("file", "gopher.json", "json w/story")
	flag.Parse()

	fmt.Printf("using story file: %v\n", *filename)
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(f)
	var story adventure.Story

	if err := decoder.Decode(&story); err != nil {
		panic(err)
	}

	fmt.Printf("%+v", story)
}
