package main

import "fmt"


func main() {
	// var card string = "Ace of Spades"
	cards := deck { newCard(), "Ace of Diamonds" }
	cards = append(cards, "Six of Spades")

	for _, card := range cards {
		fmt.Println(card)
	}
	for i, card := range cards {
		fmt.Println(i, card)
	}

	fmt.Println(cards)
}

func newCard() string {
	     return "Five of Diamonds"


}
