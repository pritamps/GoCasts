package main

import (
	"fmt"
)

func main() {

	// Declare a map - Method 1
	colors := map[string]string{
		"green": "00FF00",
		"red":   "FF0000",
		"blue":  "0000FF",
	}

	fmt.Println(colors)

	// Declare a map - Method 2
	moreColors := make(map[string]string)
	moreColors["violet"] = "aa00bb"

	printMap(colors)

}

func printMap(c map[string]string) {

	for k, v := range c {
		fmt.Println(k, v)
	}
}
