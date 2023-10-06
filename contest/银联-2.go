package main

import "sort"

func explorationSupply(station []int, pos []int) []int {
	for i, p := range pos {
		idx := sort.SearchInts(station, p)
		if idx == len(station) ||
			idx > 0 && station[idx]-p >= p-station[idx-1] {
			pos[i] = idx - 1
		} else {
			pos[i] = idx
		}
	}
	return pos
}
