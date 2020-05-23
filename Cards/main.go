package main

import (
	"fmt"
)

func main() {

	cards := newDeck()
	fmt.Println("Deck before shuffling")
	cards.print()
	fmt.Println("######################")
	cards.shuffle()
	cards.print()
	//fmt.Println(rand.Intn(5))
	//fmt.Println(rand.Intn(5))

}
