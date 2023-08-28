package main

import "testing"

func TestGame(t *testing.T) {
	player := player{hp: 100, damage: 0, armor: 0}
	wield(player)
}
