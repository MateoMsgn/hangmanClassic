package hangman

import (
	"math/rand"
	"time"
)

// fonction RandomWordIndex sert à choisir un index aléatoire

func RandomWordIndex(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}
