package main

func wateringPlants(plants []int, capacity int) int {
	steps, waterCap := 0, capacity
	for i := 0; i < len(plants); i++ {
		waterCap -= plants[i]
		if len(plants)-1 == i {
			steps++
			return steps
		}
		if plants[i+1] > waterCap {
			steps += (i*2 + 3)
			waterCap = capacity
		} else {
			steps++
		}
	}
	return steps
}
