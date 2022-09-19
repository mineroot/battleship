package player

type Token string

type Player struct {
	token Token
}

func New() *Player {
	return &Player{token: ""}
}
