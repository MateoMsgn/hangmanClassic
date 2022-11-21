package hangman

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Menu() {
	menu, error := ioutil.ReadFile("menu.txt")
	if error != nil {
		fmt.Println("The game menu can't open")
	}
	fmt.Println(string(menu))
	result := ""
	fmt.Scan(&result)
	if result == "2" {
		os.Exit(0)
	} else if result != "1" {
		fmt.Println("Number not on the menu")
	}
}
