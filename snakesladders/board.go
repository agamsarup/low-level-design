package snakesladders

import "math/rand"

type Board struct {
	size     int
	numCells int
	snakes   map[int]int
	ladders  map[int]int
}

func newBoard(size, numSnakes, numLadders int) *Board {

	board := &Board{size: size, numCells: size * size}
	board.addSnakesLadders(numSnakes, numLadders)
	return board
}

func (b *Board) getEffectivePosition(position int) int {
	if pos, ok := b.snakes[position]; ok {
		return pos
	}
	if pos, ok := b.ladders[position]; ok {
		return pos
	}

	return position
}

func (b *Board) addSnakesLadders(numSnakes, numLadders int) {

	b.snakes = make(map[int]int, numSnakes)
	b.ladders = make(map[int]int, numLadders)

	for numSnakes > 0 {
		start := 1 + rand.Intn(b.numCells-1)
		end := 1 + rand.Intn(b.numCells-1)
		if start == end {
			continue
		}
		if start < end {
			start, end = end, start
		}
		b.snakes[start] = end
		numSnakes--
	}

	for numLadders > 0 {
		start := 1 + rand.Intn(b.numCells-1)
		end := 1 + rand.Intn(b.numCells-1)
		if start == end {
			continue
		}
		if start > end {
			start, end = end, start
		}
		if _, ok := b.snakes[end]; ok {
			continue
		}
		b.ladders[start] = end
		numLadders--
	}
}
