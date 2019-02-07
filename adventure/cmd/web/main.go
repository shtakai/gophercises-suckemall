package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

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
	// NOTE: for using custom template
	//tpl := template.Must(template.New("").Parse("test")
	tpl := template.Must(template.New("").Parse(storyTmpl))

	h := adventure.NewHandler(
		story,
		adventure.WithTemplate(tpl),
		adventure.WithPathFunc(pathFn),
	)
	mux := http.NewServeMux()
	mux.Handle("/story", h)
	fmt.Printf("Starting server on port %v\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", *port), h))
}

func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]
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

var storyTmpl = `
<!doctype html>
<head>
    <meta charset="UTF-8">
    <title>Text adventure(d</title>
</head>
<body>
	<section class="page">
		<h1>{{.Title}}</h1>
		{{range .Paragraphs}}
			<p>{{.}}</p>
		{{end}}

		<ul>
			{{range .Options}}
				<li><a href="/story/{{.Arc}}">{{.Text}}</a></li>
			{{end}}
		</ul>
	</section>
    <style>
		* {
        border: 1px dotted red;
		}
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FFFCF6;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #777;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: none;
        color: #6295b5;
      }
      a:active,
      a:hover {
        color: #7792a2;
      }
      p {
        text-indent: 1em;
      }
    </style>
</body>
</html>`
