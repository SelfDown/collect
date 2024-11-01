package collect

import (
	"math/rand"
	"time"
)

var r *rand.Rand

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	r = rand.New(source)
}

func RandomInt(min, max int) int {
	return r.Intn(max-min+1) + min
}
