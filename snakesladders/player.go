package snakesladders

type Player struct {
	name     string
	position int
}

func newPlayer(name string) *Player {
	return &Player{
		name:     name,
		position: 0,
	}
}
