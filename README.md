# Greenhouse.io Go Client

[![GoDoc](https://godoc.org/github.com/escholtz/greenhouse?status.svg)](https://godoc.org/github.com/escholtz/greenhouse)
[![Go Report Card](https://goreportcard.com/badge/github.com/escholtz/greenhouse)](https://goreportcard.com/report/github.com/escholtz/greenhouse)
[![Build Status](https://api.travis-ci.org/escholtz/greenhouse.svg?branch=master)](https://travis-ci.org/escholtz/greenhouse)

Go client for the Greenhouse.io API as defined at:
https://developers.greenhouse.io/job-board.html

## Installation

```bash
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
```
