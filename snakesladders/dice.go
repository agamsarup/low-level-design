package snakesladders

import "math/rand"

type Dice struct {
	numDices int
}

func newDice(numDices int) *Dice {
	return &Dice{
		numDices: numDices,
	}
}

func (d Dice) RollDice() int {
	totalSum := 0

	for i := 0; i < d.numDices; i++ {
		totalSum += 1 + rand.Intn(6)
	}

	return totalSum
}
