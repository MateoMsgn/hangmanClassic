package hangman

// fonction Dash sert à afficher le mot avec des tirets

func Dash(word string) string {
	wordDash := []byte(word)
	for j := 0; j < len(word); j++ {
		wordDash[j] = 95
	}
	return string(wordDash)
}
