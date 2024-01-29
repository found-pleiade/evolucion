package main

import (
	"slices"
	"testing"
)

func TestGerenateFull(t *testing.T) {
	ct := CardTemplate{Card: Card{Name: "test"}, FoodSlice: []int{0, 1, 2}}
	want := []Card{{Name: "test", FoodPoints: 0}, {Name: "test", FoodPoints: 1}, {Name: "test", FoodPoints: 2}}
	got := ct.Gerenate()
	if !slices.Equal(got, want) {
		t.Errorf("Gerenate() = %v, want %v", got, want)
	}
}

func TestGerenateFoodSliceDouble(t *testing.T) {
	ct := CardTemplate{Card: Card{Name: "test"}, FoodSlice: []int{1, 1}}
	want := []Card{{Name: "test", FoodPoints: 1}, {Name: "test", FoodPoints: 1}}
	got := ct.Gerenate()
	if !slices.Equal(got, want) {
		t.Errorf("Gerenate() = %v, want %v", got, want)
	}
}

func TestGerenateFoodSliceEmpty(t *testing.T) {
	ct := CardTemplate{Card: Card{Name: "test"}, FoodSlice: []int{}}
	want := []Card{}
	got := ct.Gerenate()
	if !slices.Equal(got, want) {
		t.Errorf("Gerenate() = %v, want %v", got, want)
	}
}

func TestGerenateFoodSliceNil(t *testing.T) {
	ct := CardTemplate{Card: Card{Name: "test"}}
	want := []Card{}
	got := ct.Gerenate()
	if !slices.Equal(got, want) {
		t.Errorf("Gerenate() = %v, want %v", got, want)
	}
}
