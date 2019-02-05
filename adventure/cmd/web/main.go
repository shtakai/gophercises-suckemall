package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/shtakai/gophercises-suckemall/adventure"
)

func main() {
	filename := flag.String("file", "gopher.json", "json w/story")
	flag.Parse()

	f := openFile(filename)

	story := parseStory(f)

	game(story)
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

func parseStory(f *os.File) adventure.Story {
	decoder := json.NewDecoder(f)
	var story adventure.Story
	if err := decoder.Decode(&story); err != nil {
		panic(err)
	}
	return story
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
