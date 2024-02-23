package main

type Phase struct {
	Description  string // Description of the phase
	Name         string // Name of the phase
	Number       int    // Phase number
	NextPhase    int    // Next phase number
	IsSequential bool   // If the phase is sequential or not
}

const (
	// PhaseWait is the phase where the game is waiting for players to join.
	PhaseWait = iota
	// PhaseSelectFood is the phase where players select food cards.
	PhaseSelectFood
	// PhasePlayCards is the phase where players play cards.
	PhasePlayCards
	// PhaseRevealFood is the phase where food cards are revealed.
	PhaseRevealFood
	// PhaseActivatePriorities is the phase where priority cards are activated.
	PhaseActivatePriorities
	// PhaseFeedSpecies is the phase where species are fed.
	PhaseFeedSpecies
)

var gamePhases = map[int]Phase{
	PhaseWait:               {Description: "En attente d'autres joueurs.", IsSequential: false, Name: "Attente", Number: PhaseWait, NextPhase: PhaseSelectFood},
	PhaseSelectFood:         {Description: "Chaque joueur choisit secrètement une carte Trait de leur	main et la pose face cachée sur le plateau Point d'eau. Ce sont maintenant des cartes Nourriture végétale qui seront révélées durant la phase d'alimentation pour déterminer la quantité de nourriture végétale qui sera disponible à cette phase (le nombre de jetons Nourriture).", IsSequential: false, Name: "Sélectionner la nourriture", Number: PhaseSelectFood, NextPhase: PhasePlayCards},
	PhasePlayCards:          {Description: "Chaque joueur peut jouer autant de cartes Trait qu'il le souhaite ou les garder pour un tour suivant. Un joueur peut réaliser 1 action parmi 3 différentes pour chaque carte qu'il joue durant son tour: jouer un trait, créer une nouvelle espèce ou augmenter la taille ou la population d'une de ses espèces.", IsSequential: false, Name: "Jouer les cartes", Number: PhasePlayCards, NextPhase: PhaseRevealFood},
	PhaseRevealFood:         {Description: "Les cartes Nourriture posées sur le plateau Point d'eau sont révélées. La valeur nourriture de ces cartes est aditionnée et autant de jetons Nourriture sont placés sur le plateau Point d'eau. Si le total des cartes Nourriture est négatif, il enlève autant de jetons déjà présents sur le plateau Point d'eau (si possible).", IsSequential: false, Name: "Révéler les cartes", Number: PhaseRevealFood, NextPhase: PhaseActivatePriorities},
	PhaseActivatePriorities: {Description: "Chaque carte ayant un effet prioritaire (symbole feuille) est activée.", IsSequential: false, Name: "Activer les cartes prioritaires", Number: PhaseActivatePriorities, NextPhase: PhaseFeedSpecies},
	PhaseFeedSpecies:        {Description: "En commençant par le 1er joueur et en continuant dans le sens horaire, chaque joueur doit nourrir une de ses espèces affamées. Une espèce est affamée si elle a moins de jetons Nourriture que de Population. Quand une espèce prend de la nourriture, les jetons sont placés sur la zone prévue à cet effet au-dessus de la piste de Population en commençant par le 1er emplacement libre. Chaque jeton Nourriture rapportera un point en fin de partie.", IsSequential: true, Name: "Nourrir les espèces", Number: PhaseFeedSpecies, NextPhase: PhaseSelectFood},
}
