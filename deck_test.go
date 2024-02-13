package main

import (
	"slices"
	"testing"
)

func TestInitializeDeck(t *testing.T) {
	var want []Card
	for _, ct := range cardTemplates {
		want = append(want, ct.generate()...)
	}
	got := initializeDeck()
	if slices.Equal(got, want) {
		t.Errorf("InitializeDeck() = %v. No card shuffling seem to have happened", got)
	}
}
