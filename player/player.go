package player

type Player struct {
	name string
}

func NewPlayer(name string) Player {
	return Player{name}
}

func (p Player) ShowName() string {
	return p.name
}

func (p Player) GetMark() string {
	return "TODO: IMPLEMENT"
}
