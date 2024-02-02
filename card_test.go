package main

import (
	"slices"
	"testing"
)

func TestGenerateFull(t *testing.T) {
	ct := CardTemplate{Card: Card{Name: "test"}, FoodSlice: []int{0, 1, 2}}
	want := []Card{{Name: "test", FoodPoints: 0}, {Name: "test", FoodPoints: 1}, {Name: "test", FoodPoints: 2}}
	got := ct.Generate()
	if !slices.Equal(got, want) {
		t.Errorf("Generate() = %v, want %v", got, want)
	}
}

func TestGenerateFoodSliceDouble(t *testing.T) {
	ct := CardTemplate{Card: Card{Name: "test"}, FoodSlice: []int{1, 1}}
	want := []Card{{Name: "test", FoodPoints: 1}, {Name: "test", FoodPoints: 1}}
	got := ct.Generate()
	if !slices.Equal(got, want) {
		t.Errorf("Generate() = %v, want %v", got, want)
	}
}

func TestGenerateFoodSliceEmpty(t *testing.T) {
	ct := CardTemplate{Card: Card{Name: "test"}, FoodSlice: []int{}}
	want := []Card{}
	got := ct.Generate()
	if !slices.Equal(got, want) {
		t.Errorf("Generate() = %v, want %v", got, want)
	}
}

func TestGenerateFoodSliceNil(t *testing.T) {
	ct := CardTemplate{Card: Card{Name: "test"}}
	want := []Card{}
	got := ct.Generate()
	if !slices.Equal(got, want) {
		t.Errorf("Generate() = %v, want %v", got, want)
	}
}
