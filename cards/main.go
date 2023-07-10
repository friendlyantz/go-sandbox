package main

import "fmt"

func main() {
    // var card string = "Ace of Spades"
    // := is only used when declaring a new variable
    card_one := "Ace of Spades" 

    // = is used to reassign a variable
    card_one = "Five of Diamonds"

    fmt.Println(card_one)

    card_two := newCard()
    fmt.Println(card_two)

    cards := []string{"Joker", newCard(), card_one}

    // append does not modify the existing slice, it returns a new slice
    cards = append(cards, "Six of Spades") 

    fmt.Println(cards)

    for i, card := range cards {
        fmt.Println(i, card)
    }
}

func newCard() string { // string is the return type
    return "Eight of Hearts"
}

