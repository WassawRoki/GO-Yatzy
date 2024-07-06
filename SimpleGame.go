package main

import "strconv"

func PlaySimpleGame() {

	roundObjects := make(ScoreBoard, 13)
	// Play a simplyfied game where it simply rols 13 times a and sum up the score
	score := 0
	for i := 0; i < 13; i++ {
		diceSet := CreateDiceSet(5, 6)
		score += diceSet.Result()        // score should be roundscore which should eb evaluated by a roll manager
		roundname := strconv.Itoa(i + 1) // make the dices set the name of the round
		roundObjects[i] = &RoundResult{roundName: roundname, dices: diceSet}
		roundObjects[i].roundScore = score
	}
	roundObjects.ShowScore()
}
