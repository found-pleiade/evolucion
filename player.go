package main

type Player struct {
	ID      int
	Name    string
	Hand    []Card
	Species []Species
	IsReady bool
}

func (p *Player) toggleReady() {
	p.IsReady = !p.IsReady
}
