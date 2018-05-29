package main

import (
	"fmt"
)

func main() {
	// Declare a map
	colors := map[string]string{
		"green": "00FF00",
		"red":   "FF0000",
		"blue":  "0000FF",
	}
	fmt.Println(colors["green"])
}
