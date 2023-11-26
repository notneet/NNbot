package helper_random

import "math/rand"

func GenRand(maxNum int) int {
	// Create a new random number generator with the given seed
	randomGenerator := rand.New(rand.NewSource(int64(42)))

	return randomGenerator.Intn(maxNum)
}
