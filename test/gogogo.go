package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
)

type Game struct {
	Point int
	Round int

	Title     string
	ImagePath string
	Choice1   string
	Choice2   string
	Choice3   string
	Choice4   string
	Answer    string
}

func main() {
	Game_var := Game{
		Point:     0,
		Round:     0,
		Title:     "Title",
		ImagePath: "image/1.jpg",
		Choice1:   "Choice1",
		Choice2:   "Choice2",
		Choice3:   "Choice3",
		Choice4:   "Choice4",
		Answer:    "Answer",
	}
	writeStructToFile(Game_var, "new_index.html")

	file, err := os.Open("data/question.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	//Get a random number beetwen 0 and len(records)
	randomNumber := rand.Intn(len(records))
	Game_var.Title = records[randomNumber][0]
	Game_var.ImagePath = records[randomNumber][1]
	Game_var.Choice1 = records[randomNumber][2]
	Game_var.Choice2 = records[randomNumber][3]
	Game_var.Choice3 = records[randomNumber][4]
	Game_var.Choice4 = records[randomNumber][5]
	Game_var.Answer = records[randomNumber][6]

	writeStructToFile(Game_var, "templates/super_index.html")

	// Define a route handler
	http.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("image"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Parse and execute the HTML template from an external file
		// tmpl, err := template.ParseFiles("templates/index.html")
		tmpl, err := template.ParseFiles("templates/super_index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Render the template
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	http.HandleFunc("/handleClick", func(w http.ResponseWriter, r *http.Request) {
		// Parse the JSON request
		var requestData map[string]string
		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Handle the button click on the server side
		buttonNumber := requestData["buttonNumber"]
		message := fmt.Sprintf("Button %s clicked on the server!", buttonNumber)

		// Send a JSON response back to the client
		responseData := map[string]string{"message": message}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
	})
	// Start the web server
	http.ListenAndServe(":8080", nil)
}
