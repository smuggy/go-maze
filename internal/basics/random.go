package basics

import (
	"math/rand"
	"time"
)

var seed = time.Now().UTC().UnixNano()

func init() {
	rand.Seed(seed)
}

func ReInitRand(newSeed int64) {
	rand.Seed(newSeed)
}

func GetRand(num int) int {
	if num == 0 {
		return num
	}
	return rand.Intn(num)
}

func GetRand8(num int8) int8 {
	if num == 0 {
		return num
	}
	return int8(rand.Intn(int(num)))
}
