package services

import (
	"backend/entities"
	"math"
)

const (
	k = 30
)

func CalculateEloChange(teamAvgElo, opponentAvgElo float64, gameResult string) float64 {
	expectedScore := CalculateProbability(teamAvgElo, opponentAvgElo)

	var actualScore float64
	if gameResult == entities.ResultWin {
		actualScore = 1.0
	} else if gameResult == entities.ResultLoss {
		actualScore = 0.0
	} else {
		actualScore = 0.5
	}

	return k * (actualScore - expectedScore)
}

func CalculateProbability(eloA, eloB float64) float64 {
	return 1.0 / (1.0 + math.Pow(10, (eloB-eloA)/400.0))
}
