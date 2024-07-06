package main

import (
	"math/rand"
)

type StandardDice struct {
	Sides  int
	Result int
}

func CreateOneDice(sides int) *StandardDice {
	// Initialize a new dice with a given number of sides
	dice := &StandardDice{Sides: sides}
	dice.Roll()
	return dice
}

func (dice *StandardDice) Roll() *StandardDice {
	// Roll the dice and return the result
	dice.Result = rand.Intn(dice.Sides) + 1
	return dice
}

func (dice *StandardDice) Draw() []string {
	// Draw the dice with ASCII art
	switch dice.Result {
	case 1:
		return []string{
			"╔═════╗",
			"║     ║",
			"║  ●  ║",
			"║     ║",
			"╚═════╝",
		}
	case 2:
		return []string{
			"╔═════╗",
			"║ ●   ║",
			"║     ║",
			"║   ● ║",
			"╚═════╝",
		}
	case 3:
		return []string{
			"╔═════╗",
			"║ ●   ║",
			"║  ●  ║",
			"║   ● ║",
			"╚═════╝",
		}
	case 4:
		return []string{
			"╔═════╗",
			"║ ● ● ║",
			"║     ║",
			"║ ● ● ║",
			"╚═════╝",
		}
	case 5:
		return []string{
			"╔═════╗",
			"║ ● ● ║",
			"║  ●  ║",
			"║ ● ● ║",
			"╚═════╝",
		}
	case 6:
		return []string{
			"╔═════╗",
			"║ ● ● ║",
			"║ ● ● ║",
			"║ ● ● ║",
			"╚═════╝",
		}
	default:
		return []string{"Invalid dice result!"}
	}
}
