package rng

type RNG struct {
	seed int64
	c    int64
	m    int64
	a    int64
}

func New(seed int64) *RNG {
	return &RNG{
		seed: seed,
		c:    1234,
		m:    32767,
		a:    1103515245,
	}
}

func (rng *RNG) Float64() float64 {
	rng.seed = (rng.seed*rng.a + rng.c) & rng.m
	return float64(rng.seed) / float64(rng.m)
}
