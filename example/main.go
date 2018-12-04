package main

import (
	"fmt"
	"os"

	"github.com/escholtz/greenhouse"
)

func main() {
	client := greenhouse.NewClient()
	board, err := client.Board("github")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
		return
	}
	fmt.Printf("%+v\n", board)

	jobs, err := client.Jobs("github")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
		return
	}
	for _, j := range jobs.Jobs {
		fmt.Println(j.Title, j.Location.Name)
	}
}
