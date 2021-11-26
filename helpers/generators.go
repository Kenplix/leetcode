package helpers

import "math/rand"

// RandIntInRange returns number in range [from, to]
func RandIntInRange(min, max int) int {
	if min > max {
		panic("max value should be greater than min")
	}

	return min + rand.Intn(max-min+1)
}
