package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// bool, for JSON booleans
// float64, for JSON numbers
// string, for JSON strings
// []interface{}, for JSON arrays
// map[string]interface{}, for JSON objects
// nil for JSON null

type StorySection struct {
	Array []string `json:"story"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Story struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type StoryWrapper struct {
	Story Story `json:"intro"`
}

func parseJSON(jsonData []byte) (StoryWrapper, error) {
	var story StoryWrapper
	err := json.Unmarshal(jsonData, &story)
	if err != nil {
		return StoryWrapper{}, err
	}
	return story, nil
}

func main() {
	JsonFile, err_i := os.ReadFile("story.json")
	if err_i != nil {
		panic(err_i)
	}
	JsonData, err := parseJSON(JsonFile)
	if err != nil {
		panic(err)
	}
	// fmt.Println(JsonData)
	fmt.Print(("Title: \n"))
	fmt.Println(JsonData.Story.Title)
	fmt.Print(("Story: \n"))
	fmt.Println(JsonData.Story.Story)
	fmt.Print(("Options: \n"))
	for i := 0; i < len(JsonData.Story.Options); i++ {
		fmt.Print("Option text: \n")
		fmt.Println(JsonData.Story.Options[i].Text)
		fmt.Print("Option arc: \n")
		fmt.Println(JsonData.Story.Options[i].Arc)
	}
}
