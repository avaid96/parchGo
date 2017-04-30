package main

import (
	"math/rand"
	"time"
)

func throwDie() int {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(6) + 0
	if i == 0 {
		return 6
	}
	return i
}
