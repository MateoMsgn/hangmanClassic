package hangman

import (
	"fmt"
	"io/ioutil"
)

// fonction RandomWord sert à choisir un mot aléatoirement dans le fichier words.txt

func RandomWord() string {
	var table []string
	list, error := ioutil.ReadFile("words.txt")
	if error != nil {
		fmt.Println(error)
	}
	word := ""
	for _, value := range list {
		if value == '\n' {
			table = append(table, word)
			word = ""
		} else {
			word = word + string(value)
		}
	}
	pick := RandomWordIndex(len(table))
	return table[pick]
}
