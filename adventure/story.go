package adventure

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
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

type handler struct {
	s Story
}

func NewHandler(s Story) http.Handler {
	return handler{s}
}

func (h handler) ServeHttp(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, h.s["intro"])
	if err != nil {
		panic(err)
	}
}
