package adventure

// Parse json to map[string]Chapter

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Story map[string]Chapter

//func main() {
//	jsonStory := `
//{"intro":
//  {
//    "title": "this is title",
//    "story": [
//      "something good story",
//      "anything good story"
//    ],
//    "options": [
//      {
//        "text": "up",
//        "arc": "up up"
//      },
//      {
//        "text": "down",
//        "arc": "down down"
//      }
//    ]
//  }
//}
//
//`
//
//	//err := json.Unmarshal([]byte(jsonStory), &Story)
//	//if err != nil {
//	//	fmt.Errorf("error %v", err)
//	//}
//	//fmt.Println(Story)
//}
