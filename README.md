[![GoDoc](https://godoc.org/github.com/escholtz/greenhouse?status.svg)](https://godoc.org/github.com/escholtz/greenhouse)

Go client for the Greenhouse.io API as defined at:
https://developers.greenhouse.io/job-board.html

## Installation

```
go get -u github.com/escholtz/greenhouse
```

## Usage

```go
package main

import (
	"fmt"
	"os"

	"github.com/escholtz/greenhouse"
)

func main() {
	client := greenhouse.NewClient()
	board, err := client.Board("github")
	// handle err
	fmt.Printf("%+v\n", board)

	jobs, err := client.Jobs("github")
	// handle err
	for _, j := range jobs.Jobs {
		fmt.Println(j.Title, j.Location.Name)
	}
}
```
