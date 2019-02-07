package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"os"
	"strconv"

	"github.com/shtakai/gophercises-suckemall/adventure"
)

func main() {
	port := flag.Int("port", 3333, "the port number")
	filename := flag.String("file", "gopher.json", "json w/story")
	flag.Parse()

	f := openFile(filename)

	story, err := adventure.JsonStory(f)
	if err != nil {
		panic(err)
	}

	h := adventure.NewHandler(story, nil)
	fmt.Printf("Starting server on port %v\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", *port), h))
}

func displayFile(filename *string) (int, error) {
	return fmt.Printf("using story file: %v\n", *filename)
}

func openFile(filename *string) *os.File {
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	return f
}

func game(story adventure.Story) {
	var chapter adventure.Chapter
	chapter, ok := story[adventure.Intro]

	if !ok {
		panic("no intro")
	}

	for {
		chapter = renderChapter(chapter, story)
	}

}

func renderChapter(chapter adventure.Chapter, story adventure.Story) adventure.Chapter {
	var (
		input, arc string
	)
	renderTitle(chapter)

	for _, paragraph := range chapter.Paragraphs {
		fmt.Printf("%v\n\n", paragraph)
	}
	ensureEnd(chapter)

	renderOptions(chapter)

	fmt.Print("Input your fucked\n")
	for {
		fmt.Scan(&input)
		fmt.Printf("you selected %v\n", input)
		number, err := strconv.Atoi(input)
		if err != nil || number < 0 || number > len(chapter.Options)-1 {
			fmt.Println("fuck")
			continue
		}

		fmt.Println("ok")
		arc = chapter.Options[number].Arc
		break
	}
	fmt.Printf("arc: %v", arc)
	chapter = story[arc]
	return chapter
}

func renderOptions(chapter adventure.Chapter) {
	fmt.Print("Options\n")
	for i, option := range chapter.Options {
		fmt.Printf("%d : %v\n\n", i, option.Text)
	}
}

func ensureEnd(chapter adventure.Chapter) {
	if len(chapter.Options) == 0 {
		fmt.Println("FUCK OFFFF")
		os.Exit(0)
	}
}

func renderTitle(chapter adventure.Chapter) (int, error) {
	return fmt.Printf("TITLE: %v\n", chapter.Title)
}
