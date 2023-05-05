package grand

import (
	"math/rand"
)

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

// RandomInt func;
// Intn returns, as an int, a non-negative pseudo-random number in the half-open interval [0,n) from the default Source. It panics if n <= 0
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

//RandomNumber func
func RandomNumber(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(48, 57))
	}
	return string(bytes)
}
