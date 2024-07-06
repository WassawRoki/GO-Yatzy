package main

import "fmt"

type RoundResult struct {
	roundName  string
	dices      *DiceSet
	roundScore int
}

type ScoreBoard []*RoundResult

func (set ScoreBoard) TotalScore() int {
	score := 0
	for _, round := range set {
		score += round.dices.Result()
	}
	return score
}

func (set ScoreBoard) ShowScore() {
	board := "╔═════════════════════════════════════════════╗\n"
	for i := range set {
		roundStr := fmt.Sprintf("%-6s", set[i].roundName)
		scoreStr := fmt.Sprintf("%-6d", set[i].dices.Result())
		dicesetStr := fmt.Sprintf("%-6s", set[i].dices.RollResult())
		board += fmt.Sprintf("║ round: %-6s score: %-6s Roll: %-6s║\n", roundStr, scoreStr, dicesetStr)
	}
	board += "║═════════════════════════════════════════════║\n"
	board += fmt.Sprintf("║ TOTAL POINTS: %-30d║\n", set.TotalScore())
	board += "╚═════════════════════════════════════════════╝"
	fmt.Println(board)
}
