package main

import (
	"flag"
	"fmt"
	"goprojects/chooseyouradventure/story"
	"log"
	"net/http"
	"os"
)

func main() {

	// Flag input
	fileName := flag.String("file", "gopher.json", "File Name")
	flag.Parse()

	fmt.Printf("%s\n", *fileName)

	// Parsing the file
	jsonFile, err := os.Open(*fileName)
	if err != nil {
		log.Fatal("Could not open file")
	}
	defer jsonFile.Close()

	adventureStory, err := story.JsonToStory(jsonFile)
	if err != nil {
		log.Fatal("Could not convert json data.")
	}
	fmt.Printf("%+v\n", adventureStory)

	log.Fatal(http.ListenAndServe(":8080", story.StoryHandler(adventureStory)))
}

// d := json.NewDecoder(f)
// count := 0
// for {
// 	var chapter story.Chapter
// 	count++

// 	if err := d.Decode(&chapter); err == io.EOF {
// 		break
// 	} else if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%v\n", chapter)
// }
// fmt.Println(count)

// bufData := bufio.NewReader(f)
// data, _ := bufData.Peek(15)
// fmt.Println(string(data))
