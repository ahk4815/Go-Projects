package story

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

type Story map[string]Chapter

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

func (ch *Chapter) Display() {
	fmt.Printf("Title: %s\n", ch.Title)
	fmt.Println("Story: ", ch.Story)
	fmt.Println("Options: ", ch.Options)
}

func JsonToStory(r io.Reader) (Story, error) {
	jsonBytes, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var adventureStory Story
	json.Unmarshal(jsonBytes, &adventureStory)

	return adventureStory, nil
}

var htmlStoryTemplate = `<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
      <h1>{{.Title}}</h1>
      {{range .Story}}
        <p>{{.}}</p>
      {{end}}
      {{if .Options}}
        <ul>
        {{range .Options}}
          <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
        </ul>
      {{else}}
        <h3>The End</h3>
      {{end}}
    </body>
</html>`

func StoryHandler(adventureStory Story) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// This saves from sql injection
		t, err := template.New("test").Parse(htmlStoryTemplate)
		if err != nil {
			panic(err)
		}

		path := r.URL.Path[1:]
		fmt.Println(path)
		if path == "" {
			path = "intro"
		}
		t.Execute(w, adventureStory[path])
	}))
	return mux
}
