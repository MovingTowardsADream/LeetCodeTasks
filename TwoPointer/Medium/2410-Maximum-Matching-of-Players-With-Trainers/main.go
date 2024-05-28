package main

import "sort"

func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)
	first, second := 0, 0
	n, m := len(players), len(trainers)
	var amount int
	for first < n && second < m {
		if trainers[second] >= players[first] {
			amount++
			first++
		}
		second++
	}
	return amount
}
