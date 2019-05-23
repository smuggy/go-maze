package builders

import (
	"math/rand"
	"time"
)

var seed = time.Now().UTC().UnixNano()

func init() {
	rand.Seed(seed)
}

func reInitRand(newSeed int64) {
	rand.Seed(newSeed)
}

func getRand(num int) int {
	if num == 0 {
		return num
	}
	return rand.Intn(num)
}
