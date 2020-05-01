package supervillains

import (
	"math/rand"
	"time"
)

// All returns all superheroes from the list.go
func All() []string {
	return supervillains
}

// Random returns a single supervillain from the list.go
func Random() string {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(supervillains) - 1
	return supervillains[rand.Intn(max-min+1)+min]
}
