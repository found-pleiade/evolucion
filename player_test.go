package main

import (
	"testing"
)

func TestToggleReady(t *testing.T) {
	player := Player{IsReady: false}
	player.toggleReady()

	want := true
	got := player.IsReady
	if got != want {
		t.Errorf("toggleReady(). Got %v, want %v", got, want)
	}

	player.toggleReady()
	want = false
	got = player.IsReady
	if got != want {
		t.Errorf("toggleReady(). Got %v, want %v", got, want)
	}
}

func TestPlayerHasPriorityCard(t *testing.T) {
	player := Player{Species: []Species{{Traits: []Card{{IsPrior: false}, {IsPrior: true}, {IsPrior: false}}}}}
	want := true
	got := player.hasPriorityCard()
	if got != want {
		t.Errorf("hasPriorityCard(). Got %v, want %v", got, want)
	}

	player = Player{Species: []Species{{Traits: []Card{{IsPrior: false}, {IsPrior: false}, {IsPrior: false}}}}}
	want = false
	got = player.hasPriorityCard()
	if got != want {
		t.Errorf("hasPriorityCard(). Got %v, want %v", got, want)
	}
}

func TestAreSpeciesFed(t *testing.T) {
	player := Player{Species: []Species{{Food: 3, Population: 3}, {Food: 3, Population: 3}}}
	want := true
	got := player.areSpeciesFed()
	if got != want {
		t.Errorf("areSpeciesFed(). Got %v, want %v", got, want)
	}

	player = Player{Species: []Species{{Food: 3, Population: 3}, {Food: 3, Population: 4}}}
	want = false
	got = player.areSpeciesFed()
	if got != want {
		t.Errorf("areSpeciesFed(). Got %v, want %v", got, want)
	}
}

func TestArePlayersReady(t *testing.T) {
	players := []Player{{IsReady: true}, {IsReady: true}, {IsReady: true}}
	want := true
	got := arePlayersReady(players, false)
	if got != want {
		t.Errorf("arePlayersReady(). Got %v, want %v", got, want)
	}

	players = []Player{{IsReady: true}, {IsReady: false}, {IsReady: true}}
	want = false
	got = arePlayersReady(players, false)
	if got != want {
		t.Errorf("arePlayersReady(). Got %v, want %v", got, want)
	}
}

func TestArePlayersReadyPriority(t *testing.T) {
	players := []Player{{IsReady: false, Species: []Species{{Traits: []Card{{IsPrior: true}}}}}, {IsReady: true, Species: []Species{{Traits: []Card{{IsPrior: false}}}}}}
	want := false
	got := arePlayersReady(players, true)
	if got != want {
		t.Errorf("arePlayersReady(). Got %v, want %v", got, want)
	}

	players = []Player{{IsReady: true, Species: []Species{{Traits: []Card{{IsPrior: false}}}}}, {IsReady: true, Species: []Species{{Traits: []Card{{IsPrior: false}}}}}}
	want = true
	got = arePlayersReady(players, true)
	if got != want {
		t.Errorf("arePlayersReady(). Got %v, want %v", got, want)
	}
}
