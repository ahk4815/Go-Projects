package main

import (
	"flag"
	"fmt"
	htmlparser "goprojects/htmlparser/htmlParser"
	"log"
	"os"
)

func main() {
	templateDir := "templates/"
	fileName := flag.String("filePath", "ex1.html", "Path of the html file")
	flag.Parse()

	file, err := os.Open(templateDir + *fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	links := htmlparser.BuildLink(file)
	fmt.Println(links)
}
