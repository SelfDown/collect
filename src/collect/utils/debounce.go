package collect

import (
	"time"
)

func Debounce(fn func(), ms float64) func() {
	prev := time.Unix(0, 0)
	return func() {
		curr := time.Now()
		delta := curr.Sub(prev).Seconds()
		if delta < ms {
			return
		}
		prev = curr
		fn()
	}
}
