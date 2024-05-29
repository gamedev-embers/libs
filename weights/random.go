package weights

import "math/rand"

type Random struct {
	*rand.Rand
}

// non-goroutine-safe
func NewRandom(seed int64) *Random {
	s := rand.NewSource(seed)
	return &Random{
		Rand: rand.New(s),
	}
}
