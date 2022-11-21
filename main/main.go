package main

import (
	"fmt"

	"hangman"
)

func main() {
	hangman.Menu()
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	word := hangman.RandomWord()
	dashes := hangman.Dash(word)
	fmt.Println("\033[36mGood Luck\033[0m, \033[32myou have 10 attempts\033[0m.")
	wrongAttempt := 10
	usedletter := ""
	count := 0
	for attempt := 10; attempt > 0 && dashes != word; {
		fmt.Println(dashes)
		fmt.Println("Your letter is \033[31malready used\033[0m or \033[31mnot correct\033[0m :", usedletter)
		fmt.Print("Choose : ")
		guess := hangman.Input()
		if word == guess {
			dashes = guess
		} else if hangman.CorrectOrFalse(word, guess) {
			dashes = hangman.RevealWord(word, dashes, guess)
		} else {
			usedletter += guess
			attempt = attempt - 1
			count = count + 1
			fmt.Printf("\033[36mNot in the word, %v attempts remaining\033[0m \n", attempt)
			hangman.Draw(count)
			wrongAttempt = attempt
		}
	}
	if wrongAttempt == 0 {
		fmt.Printf(colorRed+"Sorry, the word was \033[32m%v\033[0m \n", word)
	} else {
		fmt.Printf(colorGreen+"\033[32mCongrats ! You win the word was %v\033[0m \n", word)
	}
}
