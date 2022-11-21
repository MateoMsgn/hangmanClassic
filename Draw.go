package hangman

import (
	"bufio"
	"fmt"
	"os"
)

// fonction Draw sert à dessiner le pendu

func Draw(count int) {
	Index := (count - 1) * 8
	draw, error := os.Open("draw.txt")
	if error != nil {
		fmt.Println(error)
	}
	file := bufio.NewScanner(draw)
	file.Split(bufio.ScanLines)
	var line []string
	for file.Scan() {
		line = append(line, file.Text())
	}
	for i := 0; i < 8; i++ {
		fmt.Println(string(line[Index+i]))
	}

}
