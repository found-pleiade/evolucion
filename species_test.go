package main

import (
	"testing"
)

func TestHasPriorityCard(t *testing.T) {
	species := Species{Traits: []Card{{IsPrior: false}, {IsPrior: true}, {IsPrior: false}}}
	want := true
	got := species.hasPriorityCard()
	if got != want {
		t.Errorf("hasPriorityCard(). Got %v, want %v", got, want)
	}

	species = Species{Traits: []Card{{IsPrior: false}, {IsPrior: false}, {IsPrior: false}}}
	want = false
	got = species.hasPriorityCard()
	if got != want {
		t.Errorf("hasPriorityCard(). Got %v, want %v", got, want)
	}
}

func TestIsFed(t *testing.T) {
	species := Species{Food: 3, Population: 3}
	want := true
	got := species.isFed()
	if got != want {
		t.Errorf("isFed(). Got %v, want %v", got, want)
	}

	species = Species{Food: 3, Population: 4}
	want = false
	got = species.isFed()
	if got != want {
		t.Errorf("isFed(). Got %v, want %v", got, want)
	}
}
