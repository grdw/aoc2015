package main

import (
	"fmt"
	"math"
)

type player struct {
	hp     int
	damage int
	armor  int
}

type item struct {
	t      string
	cost   int
	damage int
	armor  int
}

type combo struct {
	of     []string
	cost   int
	damage int
	armor  int
}

func main() {
	fmt.Println("part 1:", part1())
}

func part1() int {
	items := []item{
		{"weapon", 8, 4, 0},
		{"weapon", 10, 5, 0},
		{"weapon", 25, 6, 0},
		{"weapon", 40, 7, 0},
		{"weapon", 74, 8, 0},
		{"armor", 13, 0, 1},
		{"armor", 31, 0, 2},
		{"armor", 53, 0, 3},
		{"armor", 75, 0, 4},
		{"armor", 102, 0, 5},
		{"ring", 25, 1, 0},
		{"ring", 50, 2, 0},
		{"ring", 100, 3, 0},
		{"ring", 25, 0, 1},
		{"ring", 50, 0, 2},
		{"ring", 100, 0, 3},
	}
	combos := inventoryCombos(items)
	fmt.Println(combos)

	minCost := math.MaxInt32
	for _, combo := range combos {
		you := player{hp: 100, damage: 0, armor: 0}
		cost := wield(you, combo)
		boss := player{hp: 100, damage: 8, armor: 2}
		win := game(you, boss)
		if win && cost < minCost {
			minCost = cost
		}
	}
	return minCost
}

func inventoryCombos(items []item) []combo {
	totals := []combo{}
	for _, i := range items {
		list := []string{i.t}
		iOnly := combo{list, i.cost, i.damage, i.armor}

		//for _, a := range armors {
		//	wPlusA := item{
		//		"combo",
		//		w.cost + a.cost,
		//		w.damage + a.damage,
		//		w.armor + a.armor,
		//	}

		//	totals = append(totals, wPlusA)
		//}
		totals = append(totals, iOnly)
	}
	return totals
}

func wield(p player, c combo) int {
	p.damage = c.damage
	p.armor = c.armor
	return c.cost
}

func game(p1 player, p2 player) bool {
	turn := "p1"
	for p1.hp <= 0 || p2.hp <= 0 {
		if turn == "p1" {
			damage := p1.damage - p2.armor
			if damage < 0 {
				damage = 1
			}
			p2.hp -= damage
			turn = "p2"
		} else if turn == "p2" {
			damage := p2.damage - p1.armor
			if damage < 0 {
				damage = 1
			}
			p1.hp -= damage
			turn = "p1"
		}
		fmt.Println(turn, p1, p2)
	}
	return p2.hp <= 0
}
