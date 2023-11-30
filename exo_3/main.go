package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

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
	Story Story `json:"-"`
}

func parseJSON(jsonData []byte, key string) (Story, error) {
	var storyMap map[string]Story
	err := json.Unmarshal(jsonData, &storyMap)
	if err != nil {
		return Story{}, err
	}

	story, ok := storyMap[key]
	if !ok {
		return Story{}, fmt.Errorf("Key '%s' not found in JSON", key)
	}

	return story, nil
}

func printStory(story Story) {
	fmt.Print("Title: \n")
	fmt.Println(story.Title)
	fmt.Print("Story: \n")
	fmt.Println(story.Story)
	fmt.Print("Options: \n")
	for i := 0; i < len(story.Options); i++ {
		fmt.Print("Option text: \n")
		fmt.Println(story.Options[i].Text)
		fmt.Print("Option arc: \n")
		fmt.Println(story.Options[i].Arc)
	}
	fmt.Print("\n\n\n\n\n\n\n\n")
}

func main() {
	JsonFile, err_i := os.ReadFile("story.json")
	if err_i != nil {
		panic(err_i)
	}
	arc := "intro"
	JsonData, err := parseJSON(JsonFile, arc)
	if err != nil {
		panic(err)
	}
	printStory(JsonData)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>%s</h1>", JsonData.Title)
		fmt.Fprintf(w, "<p>%s</p>", JsonData.Story)
		for i := 0; i < len(JsonData.Options); i++ {
			for i := 0; i < len(JsonData.Options); i++ {
				fmt.Fprintf(w, "<a href='/choose/%d'><p>%s</p></a>", i, JsonData.Options[i].Text)
			}
		}

		http.HandleFunc("/choose/", func(w http.ResponseWriter, r *http.Request) {
			// optionIndex := extractOptionIndex(r.URL.Path)
			var optionIndex int
			optionIndex = 0
			if optionIndex >= 0 && optionIndex < len(JsonData.Options) {
				selectedOption := JsonData.Options[optionIndex]
				JsonData, err = parseJSON(JsonFile, selectedOption.Arc)
			}
		})
	})
	http.ListenAndServe(":8080", nil)
}
