package main

type Player struct {
	ID           int
	Name         string
	Hand         []Card
	Species      []Species
	IsReady      bool
	SelectedCard int
}

// initialize initializes the player.
func (p *Player) initialize(id int, name string) {
	p.SelectedCard = -1
}

// toggleReady toggles the player's ready status.
func (p *Player) toggleReady() {
	p.IsReady = !p.IsReady
}

// hasPriorityCard checks if the player has a priority card on one of his species.
func (p *Player) hasPriorityCard() bool {
	for _, species := range p.Species {
		if species.hasPriorityCard() {
			return true
		}
	}
	return false
}

// areSpeciesFed checks if all species are fed.
func (p *Player) areSpeciesFed() bool {
	for _, species := range p.Species {
		if !species.isFed() {
			return false
		}
	}
	return true
}

// arePlayersReady checks if players are ready.
// If priority is true, it will only check status of players with priority cards equipped.
func arePlayersReady(players []Player, priority bool) bool {
	for _, player := range players {
		if priority && !player.hasPriorityCard() {
			continue
		}
		if !player.IsReady {
			return false
		}
	}
	return true
}
