package main

import "fmt"


func main() {
	// var card string = "Ace of Spades"
	cards := []string { newCard(), "Ace of Diamonds" }
	cards = append(cards, "Six of Spades")

	fmt.Println(cards)
}

func newCard() string {
	     return "Five of Diamonds"


}
