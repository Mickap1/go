package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
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

type StoryHandler struct {
	Story_key     []string
	StoryLink     []string
	Story         []Story
	Nb_Story_Page int
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

func StoryHandlerFunc(storyMap map[string]Story) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Path[len("/story/"):]
		story, ok := storyMap[key]
		if !ok {
			http.Error(w, "Story not found.", http.StatusNotFound)
			return
		}
		json, err := json.Marshal(story)
		if err != nil {
			http.Error(w, "Unable to marshal JSON.", http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(json)
	}
}

func fill_StoryHandler(JsonData Story, StoryHandler *StoryHandler, JsonFile []byte) {
	var boolean bool
	boolean = false
	for i := 0; i < len(JsonData.Options); i++ {
		fmt.Println(i)
		boolean = true
		for j := 0; j < StoryHandler.Nb_Story_Page; j++ {
			if JsonData.Options[i].Arc == StoryHandler.Story_key[j] {
				fmt.Printf("%s already exist %s\n", JsonData.Options[i].Arc, StoryHandler.Story_key[j])
				boolean = false
			}
		}
		if boolean {
			StoryHandler.Nb_Story_Page++
			StoryHandler.Story_key[StoryHandler.Nb_Story_Page-1] = JsonData.Options[i].Arc
			StoryHandler.StoryLink[StoryHandler.Nb_Story_Page-1] = "/story/" + JsonData.Options[i].Arc
			nJsonFile, err_i := parseJSON(JsonFile, JsonData.Options[i].Arc)
			if err_i != nil {
				fmt.Println("NOOOOO THERE IS PANIC")
				panic(err_i)
			}
			StoryHandler.Story[StoryHandler.Nb_Story_Page-1] = nJsonFile
			fill_StoryHandler(nJsonFile, StoryHandler, JsonFile)
		}
	}
}

func get_choose_index(StoryHandler StoryHandler, title string, increment int) int {
	fmt.Println("------------------------")
	fmt.Println(title)
	fmt.Println(increment)
	var var1 int
	var1 = -1
	var var2 string
	var2 = "No"
	var index int
	index = -1
	for i := 0; i < StoryHandler.Nb_Story_Page; i++ {
		if StoryHandler.Story_key[i] == title {
			var1 = i
		}
	}
	var2 = StoryHandler.Story[var1].Options[increment].Arc
	for j := 0; j < StoryHandler.Nb_Story_Page; j++ {
		if StoryHandler.Story_key[j] == var2 {
			index = j
		}
	}
	return index
}

func handle_choice(StoryHandler StoryHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		indexStr := r.URL.Path[len("/choose/"):]
		choose_index, err := strconv.Atoi(indexStr)
		if err != nil {
			http.Error(w, "Invalid index", http.StatusBadRequest)
			return
		}
		if choose_index < 0 || choose_index >= StoryHandler.Nb_Story_Page {
			http.Error(w, "Invalid index", http.StatusBadRequest)
			return
		}
		fmt.Println(StoryHandler.Story_key[choose_index])
		fmt.Println(StoryHandler.StoryLink[choose_index])
		fmt.Fprintf(w, "<h1>%s</h1>", StoryHandler.Story[choose_index].Title)
		fmt.Fprintf(w, "<p>%s</p>", StoryHandler.Story[choose_index].Story)
		for i := 0; i < len(StoryHandler.Story[choose_index].Options); i++ {
			fmt.Fprintf(w, "<a href='/choose/%d'><p>%s</p></a>", get_choose_index(StoryHandler, StoryHandler.Story_key[choose_index], i), StoryHandler.Story[choose_index].Options[i].Text)
		}
	}
}

func main() {
	var StoryHandler StoryHandler
	JsonFile, err_i := os.ReadFile("story.json")
	if err_i != nil {
		panic(err_i)
	}
	arc := "intro"
	JsonData, err := parseJSON(JsonFile, arc)
	StoryHandler.Story = make([]Story, 100)
	StoryHandler.Story_key = make([]string, 100)
	StoryHandler.StoryLink = make([]string, 100)
	StoryHandler.Story[0] = JsonData
	StoryHandler.Story_key[0] = arc
	StoryHandler.StoryLink[0] = "/story/" + arc
	StoryHandler.Nb_Story_Page = 0
	if err != nil {
		panic(err)
	}
	fill_StoryHandler(JsonData, &StoryHandler, JsonFile)
	for i := 0; i < StoryHandler.Nb_Story_Page; i++ {
		fmt.Println("--------------------")
		fmt.Println(i)
		fmt.Println(StoryHandler.Story_key[i])
		fmt.Println(StoryHandler.StoryLink[i])
		fmt.Println("--------------------")
		fmt.Println()
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>%s</h1>", JsonData.Title)
		fmt.Fprintf(w, "<p>%s</p>", JsonData.Story)
		for i := 0; i < len(JsonData.Options); i++ {
			for i := 0; i < len(JsonData.Options); i++ {
				fmt.Fprintf(w, "<a href='/choose/%d'><p>%s</p></a>", i, JsonData.Options[i].Text)
			}
		}
	})
	http.HandleFunc("/choose/", handle_choice(StoryHandler))
	http.ListenAndServe(":8080", nil)
}
