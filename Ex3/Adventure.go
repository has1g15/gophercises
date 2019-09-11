package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	parseStory("story.json")
}

func parseStory(fileName string) {
	jsonFile, err := os.Open(fileName)
	content, _ := ioutil.ReadAll(jsonFile)

	if err != nil {
		log.Println(err)
	}

	var sections Sections

	json.Unmarshal(content, &sections)

	fmt.Println(sections.Sections[0])

	for i := 0; i < len(sections.Sections); i++ {
		fmt.Println("Section Name: " + sections.Sections[i].name)
		fmt.Println("Section Title: " + sections.Sections[i].title)
		fmt.Println("Story: " + sections.Sections[i].story)
		fmt.Println("Options: " + sections.Sections[i].options.text)
	}
}

type Sections struct {
	Sections[] Section
}

type Section struct {
	name string
	title string
	story string
	options Options
}

type Options struct {
	text string
	arc string
}

