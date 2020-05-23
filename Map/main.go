package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "656ff",
		"green": "6y8878",
		"Blue":  "67767",
	}

	printMap(colors)

}

func printMap(c map[string]string) {

	for color, hex := range c {
		fmt.Println("The code for ", color, "is ", hex)
	}
}
