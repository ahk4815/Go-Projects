package story

import "fmt"

type Story map[string]Chapter

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"chapter"`
}

func (ch *Chapter) Display() {
	fmt.Printf("Title: %s\n", ch.Title)
	fmt.Println("Story: ", ch.Story)
	fmt.Println("Options: ", ch.Options)
}
