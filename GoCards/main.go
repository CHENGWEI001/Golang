package main

// import "fmt"

func main() {
	cards := newDeck()
	// 	dealCards, _ := deal(cards, 5)
	// 	//fmt.Printf("%p, %p, %p\n", &cards[0], &dealCards[0], &remainingCards[0])
	// 	fmt.Printf("ori %s %s\n", cards[0], dealCards[0])
	// 	cards[0] = "A"
	// 	fmt.Printf("after %s %s\n", cards[0], dealCards[0])
	// 	//cards.saveToFile("my_cards")
	// 	//cards := newDecFromFile("my_cards")
	cards.shuffle()
	cards.print()
}
