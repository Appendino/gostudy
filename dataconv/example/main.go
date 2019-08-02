package main

import "github.com/Appendino/gostudy/dataconv"

func main() {
	if err := dataconv.Showconv(); err != nil {
		panic(err)
	}

	dataconv.Interfaces()
}
