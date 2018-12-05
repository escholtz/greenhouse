package main

import (
	"fmt"

	"github.com/escholtz/greenhouse"
)

func main() {
	client := greenhouse.NewClient()

	// Get the company name
	board, err := client.Board("github")
	if err != nil {
		panic(err)
	}
	fmt.Println(board.Name)

	// Get all job openings including descriptions
	jobs, err := client.Jobs("github")
	if err != nil {
		panic(err)
	}
	for _, j := range jobs.Jobs {
		fmt.Println(j.Title, j.Location.Name)
	}
}
