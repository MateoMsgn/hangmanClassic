package hangman

// fonction RevealWord sert à afficher le mot avec les lettres trouvées

func RevealWord(word string, dashes string, guess string) string {
	guessByte := []byte(guess)
	dash_Byte := []byte(dashes)
	for i := 0; i < len(word); i++ {
		if word[i] == guess[0] {
			dash_Byte[i] = guessByte[0]
		}
	}
	return string(dash_Byte)
}
