package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Parse json to map[string]Chapter

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func main() {
	jsonStory := `
  {
    "title": "this is title",
    "story": [
      "something good story",
      "anything good story"
    ],
    "options": [
      {
        "text": "up",
        "arc": "up up"
      },
      {
        "text": "down",
        "arc": "down down"
      }
    ]
  }
`

	var Story Chapter
	_ = Story

	decoder := json.NewDecoder(strings.NewReader(jsonStory))
	err := decoder.Decode(&Story)
	if err != nil {
		fmt.Errorf("error %v", err)
	}
	fmt.Println(Story)
}
