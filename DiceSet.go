package main

import "fmt"

type DiceSet struct {
	Dices []*StandardDice
}

func CreateDiceSet(numDices, sides int) *DiceSet {
	// Initialize a new set of dices with a given number of dices and sides
	dices := make([]*StandardDice, numDices)
	for i := range dices {
		dices[i] = CreateOneDice(sides)
	}
	return &DiceSet{Dices: dices}
}

func (set *DiceSet) Draw() {
	// Draw the set of dices with ASCII art on the same line
	combinedDrawing := make([]string, 5)

	// Combine each dice's drawing line by line
	for _, dice := range set.Dices {
		diceDrawing := dice.Draw()
		for i := range combinedDrawing {
			combinedDrawing[i] += diceDrawing[i] + " "
		}
	}

	// Combine the lines into a single string
	result := ""
	for _, line := range combinedDrawing {
		result += line + "\n"
	}
	fmt.Println(result)
}

func (set *DiceSet) Roll() {
	// Roll all the dices in the set
	for _, dice := range set.Dices {
		dice.Roll()
	}
}
func (set *DiceSet) Result() int {
	// Calculate the total score of the dice set
	score := 0
	for _, dice := range set.Dices {
		score += dice.Result
	}
	return score
}

func (set *DiceSet) RollResult() string {
	diceResults := ""
	for _, dice := range set.Dices {
		diceResults += fmt.Sprintf("║%d", dice.Result)
	}
	return diceResults
}
