package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	var time_limit int = 30
	var quiz_file string = "problems.csv"
	var shuffle bool = false
	var count int = 0

	flag.IntVar(&time_limit, "time", 30, "Time limit for the quiz")
	flag.StringVar(&quiz_file, "scv", "problems.csv", "Quiz file for the quiz")
	flag.BoolVar(&shuffle, "shuffle", false, "Shuffle the quiz")

	flag.Parse()

	var deadline = (time.Duration(time_limit) * time.Second)

	fmt.Printf("Time limit: %v\n", time_limit)
	fmt.Printf("Quiz file: %s\n", quiz_file)
	fmt.Printf("Shuffle: %v\n", shuffle)

	file, err := os.Open(quiz_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Shuffle the quiz
	if shuffle {
		for i := range records {
			j := i + rand.Intn(len(records)-i)
			records[i], records[j] = records[j], records[i]
		}
	}

	for _, record := range records {
		fmt.Printf("Question N°%v: %s\n", count, record[0])
		c := make(chan string, 1)
		go scan(c)
		ctx, _ := context.WithTimeout(context.Background(), deadline)
		select {
		case <-ctx.Done():
			fmt.Println("Timeout: You didn't type within the deadline.")
			fmt.Printf("You scored %v out of %v\n", count, len(records))
			os.Exit(0)
		case userInput := <-c:
			if userInput == record[1] {
				fmt.Println("Correct")
			} else {
				fmt.Println("Incorrect")
				fmt.Printf("You scored %v out of %v\n", count, len(records))
				os.Exit(0)
			}
		}
		count++
	}

}
func scan(in chan string) {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err)
	}
	in <- input
}
