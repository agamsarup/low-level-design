package snakesladders

import (
	"fmt"
	"strconv"
)

type Game struct {
	board          *Board
	dice           *Dice
	players        []*Player
	playerWithTurn int
	winner         *Player
}

func newGame(numPlayers int) *Game {
	game := &Game{
		board: newBoard(10, 5, 5),
		dice:  newDice(1),
	}

	game.addPlayers(numPlayers)

	return game
}

func (g *Game) addPlayers(numPlayers int) {
	for i := 0; i < numPlayers; i++ {
		g.players = append(g.players, newPlayer("P"+strconv.Itoa(i+1)))
	}
}

func (g *Game) whoToMoveNext() *Player {
	player := g.players[g.playerWithTurn]
	return player
}

func (g *Game) updateTurn() {
	g.playerWithTurn = (g.playerWithTurn + 1) % len(g.players)
}

func (g *Game) Play() {
	numMoves := 1000
	for g.winner == nil || numMoves < 1000 {
		nextMover := g.whoToMoveNext()
		jump := g.dice.RollDice()
		curPosition := nextMover.position
		nextPosition := curPosition + jump
		if nextPosition >= g.board.numCells {
			g.winner = nextMover
			break
		}
		nextMover.position = g.board.getEffectivePosition(nextPosition)
		if nextMover.position > nextPosition {
			fmt.Printf("Player %s used ladder from %d to %d\n", nextMover.name, nextPosition, nextMover.position)
		}
		if nextMover.position < nextPosition {
			fmt.Printf("Player %s was bitten by snake from %d to %d\n", nextMover.name, nextPosition, nextMover.position)
		}
		g.printMove(nextMover, curPosition, nextMover.position)
		g.updateTurn()
		numMoves++
	}
	g.announceWinner()
}

func (g *Game) printMove(player *Player, prev, cur int) {
	fmt.Printf("Player %s moves from %d to %d\n", player.name, prev, cur)
}

func (g *Game) announceWinner() {
	fmt.Printf("\nWINNER IS %s\n", g.winner.name)
}
