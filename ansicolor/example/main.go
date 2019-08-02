package main

import (
	"fmt"
	"github.com/Appendino/gostudy/ansicolor"
)

func main() {
	r := ansicolor.ColorText{
		TextColor: ansicolor.Red,
		Text:      "I'm red",
	}

	fmt.Println(r.String())
}
