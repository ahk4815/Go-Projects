package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"goprojects/chooseyouradventure/story"
	"io"
	"log"
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

	jsonBytes, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("Could not convert json file into byte array")
	}

	var adventureStory story.Story
	json.Unmarshal(jsonBytes, &adventureStory)

	for name, chapter := range adventureStory {
		fmt.Println(name)
		chapter.Display()
		break
	}
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
