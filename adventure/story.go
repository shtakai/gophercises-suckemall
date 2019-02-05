package adventure

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

// Parse json to map[string]Chapter
func JsonStory(r io.Reader) (Story, error) {
	decoder := json.NewDecoder(r)
	var story Story
	if err := decoder.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Story map[string]Chapter

const Intro = "intro"

var defaultHandlerTemplate = `
<!doctype html>
<head>
    <meta charset="UTF-8">
    <title>Text adventure(d</title>
</head>
<body>
    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
        <p>{{.}}</p>
    {{end}}

    <ul>
        {{range .Options}}
            <li><a href="{{.Arc}}">{{.Text}}</a></li>
        {{end}}
    </ul>
</body>
</html>`

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTemplate))
}

//type handler struct {
//	s      Story
//	t      *template.Template
//	pathFn func(r *http.Request) string
//}

func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	// '/intro' => 'intro'
	path = path[1:]

	if chapter, ok := h.s[path]; ok {
		err := tpl.Execute(w, chapter)
		if err != nil {
			log.Printf("%v\n", err)
			http.Error(w, "something went wrong", http.StatusInternalServerError)

		}
		return
	}
	http.Error(w, "chapter not found", http.StatusNotFound)

}
