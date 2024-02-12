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
