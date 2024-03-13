package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	filename, timeLimit := readArguments()
	records := readCsvFile(filename)
	counter := 0

	fmt.Println("Welcome to the Quiz Game!")
	fmt.Println("Press Enter to start the quiz.")
	fmt.Scanln()
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	go func() {
		<-timer.C
		fmt.Println("\nTime's up!")
		fmt.Printf("You scored %d out of %d\n", counter, len(records))
		os.Exit(0)
	}()

	for _, v := range records {
		q, a := v[0], v[1]
		fmt.Printf("Question: %s = ?\n", q)
		var ans string
		fmt.Scanln(&ans)
		if ans == a {
			counter++
		}

	}
	fmt.Printf("You scored %d out of %d\n", counter, len(records))
}
func readArguments() (string, int) {
	filename := flag.String("filename", "problems.csv", "CSV File that contains quiz questions")
	limit := flag.Int("limit", 5, "Time limit for the quiz in seconds")
	flag.Parse()
	return *filename, *limit
}
func readCsvFile(fileName string) [][]string {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	recordes, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return recordes
}
