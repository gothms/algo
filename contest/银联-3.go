package main

func storedEnergy(storeLimit int, power []int, supply [][]int) int {
	minVal := func(a, b int) int {
		if b < a {
			return b
		}
		return a
	}
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	store, right, min, max, n := 0, 0, 0, 0, len(supply)
	for i, p := range power {
		if right < n && i == supply[right][0] {
			min, max = supply[right][1], supply[right][2]
			right++
		}
		if p > max {
			store = minVal(storeLimit, store+p-max)
		} else if p < min {
			store = maxVal(0, store+p-min)
		}
	}
	return store
}
