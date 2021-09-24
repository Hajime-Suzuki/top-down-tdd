package players

type Player struct {
	name string
	mark string
}

func NewPlayer(name string, mark string) Player {
	return Player{name, mark}
}

func (p Player) ShowName() string {
	return p.name
}

func (p Player) GetMark() string {
	return p.mark
}
