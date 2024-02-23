package main

type Species struct {
	BodySize   int
	Food       int
	Population int
	Name       string
	Traits     []Card
}

// hasPriorityCard checks if the species has a priority card.
func (s *Species) hasPriorityCard() bool {
	for _, trait := range s.Traits {
		if trait.IsPrior {
			return true
		}
	}
	return false
}

// isFed checks if the species is fed.
func (s *Species) isFed() bool {
	return s.Food >= s.Population
}
