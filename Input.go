package hangman

import "fmt"

// fonction Input sert à récupérer la lettre entrée par l'utilisateur

func Input() string {
	var guess string
	fmt.Scanln(&guess)
	return guess
}
