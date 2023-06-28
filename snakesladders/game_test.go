package snakesladders

import "testing"

func TestGamePlay(t *testing.T) {
	game := newGame(2)
	game.Play()
}
