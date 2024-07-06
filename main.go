package main

import (
	"fmt"
)

func main() {
	var userinput string
	fmt.Println("Enter the gamemode you want to play: 1. SimpleGame 2. ManualGame 3. AdvancedGame")
	fmt.Scanln(&userinput)
	if userinput == "1" {
		PlaySimpleGame()
	}
	if userinput == "2" {
		fmt.Println("Not implemented yet")
		// PlayManualGame()
	}
	if userinput == "3" {
		fmt.Println("Not implemented yet")
		// PlayAdvancedGame()
	}

}
