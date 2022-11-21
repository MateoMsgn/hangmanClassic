package hangman

// fonction CorrectOrFalse sert à vérifier si la lettre entrée par l'utilisateur est dans le mot

func CorrectOrFalse(word string, guess string) bool {
	for _, v := range word {
		if string(v) == guess {
			return true
		}
	}
	return false
}
