package main

type CardTemplate struct {
	Card      Card
	FoodSlice []int
}

var (
	carnivoreTemplate     = CardTemplate{Card: Card{Name: "Carnivore", ShortDescription: "Doit attaquer d'autres espèces pour se nourrir et ne peut jamais manger de Plante.", LongDescription: "Cette espèce peut attaquer d'autres espèces lors de la phase Nourriture si sa Taille est supérieure à la Taille de l'espèce attaquée et si elle dispose des traits nécessaires pour contrecarrer les traits défensifs de celle-ci. Les Carnivores ne peuvent jamais manger de Plantes, même s'ils possèdent des cartes Trait comme Long Cou, Coopération ou Fourrageuse.", Color: "danger", Carnivore: true, IsPrior: false}, FoodSlice: []int{0, 0, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 6, 6}}
	carapaceTemplate      = CardTemplate{Card: Card{Name: "Carapace", ShortDescription: "+4 en Taille lorsqu'elle est attaquée.", LongDescription: "La Taille de cette espèce est augmentée de 4 lorsqu'elle est attaquée par un Carnivore. Une espèce de Taille 6 ayant une Carapace aura une Taille effective de 10. Carapace n'augmente pas la quantité de Viande qu'un Carnivore prend après l'attaque.", Color: "white", Carnivore: false, IsPrior: false}, FoodSlice: []int{1, 2, 3, 3, 4, 4, 5}}
	charognardTemplate    = CardTemplate{Card: Card{Name: "Charognard", ShortDescription: "Prend 1 Viande dans la Réserve à charque fois qu'un Carnivore attaque une espèce.", LongDescription: "À chaque fois que la Population d'une espèce est réduite suite à l'attaque d'un Carnivore, prenez 1 Viande dans la Réserve. La carte Trait Charognard s'active même s'il est porté par le Carnivore attaquant ou par l'espèce attaquée (elle prend 1 Viande, mais seulement après avoir baissé sa Population de 1).", Color: "success", Carnivore: false, IsPrior: false}, FoodSlice: []int{2, 3, 4, 5, 6, 6, 7}}
	chasseEnMeuteTemplate = CardTemplate{Card: Card{Name: "Chasse en Meute", ShortDescription: "Lorsqu'elle attaque, sa Taille estégale à la somme de sa Population et de sa Taille.", LongDescription: "Lorsqu'elle attaque, la Taille de cette espèce est égale à la somme de sa Population et de sa Taille. Si elle a une Population de 5 et une Taille de 3, elle disposera d'une Taille effective de 8 pour attaquer d'autres espèces.", Color: "danger", Carnivore: true, IsPrior: false}, FoodSlice: []int{-3, -2, -1, 0, 1, 2, 3}}
	cooperationTemplate   = CardTemplate{Card: Card{Name: "Coopération", ShortDescription: "Lorsqu'elle prend de la nourriture votre espèce à sa droite prend 1 Nourriture de la même source.", LongDescription: "À chaque fois que cette espèce prend de la nourriture, votre espèce placée à sa droite prend 1 Nourriture du même type (Plante ou Viande) et de la même source (Plan d'eau ou Réserve). Coopération sera donc activée par des cartes Trait comme Long Cou, Charognard, Intelligence, ou par Coopération (mais pas par la carte Tissus Graisseux). Exemple : si une espèce a Coopération et qu'elle prend 2 Nourritures dans la Réserve, l'espèce à sa droite prend également 1 Nourriture dans la Réserve.", Color: "success", Carnivore: false, IsPrior: false}, FoodSlice: []int{0, 3, 3, 4, 4, 5, 5}}
	cornesTemplate        = CardTemplate{Card: Card{Name: "Cornes", ShortDescription: "La Population d'un Cornivore qui l'attaque est réduite de 1.", LongDescription: "Lorsqu'un Carnivore attaque cette espèce, la Population du Carnivore est réduite de 1. Cette perte de Population est subie avant que le Carnivore ne récupère la Viande de cette attaque.", Color: "white", Carnivore: false, IsPrior: false}, FoodSlice: []int{1, 2, 3, 3, 4, 4, 5}}
	longCouTemplate       = CardTemplate{Card: Card{Name: "Long cou", ShortDescription: "Avant que les cartes Nourriture soient révélées, elle prend 1 Plante de la Réserve.", LongDescription: "Avant que les cartes Nourriture ne soient révélées, elle prend 1 Plante de la Réserve.", Color: "success", Carnivore: false, IsPrior: true}, FoodSlice: []int{3, 4, 5, 6, 7, 8, 9}}

	cardTemplates = []CardTemplate{carnivoreTemplate, carapaceTemplate, charognardTemplate, chasseEnMeuteTemplate, cooperationTemplate, cornesTemplate, longCouTemplate}
)

func (ct *CardTemplate) generate() []Card {
	var cards []Card
	for _, food := range ct.FoodSlice {
		card := ct.Card
		card.FoodPoints = food
		cards = append(cards, card)
	}
	return cards
}
