package limitlessled

import (
	"math"
)

type Repeat struct {
	int
}

func (repeat Repeat) Times(do func()) {
	for i := 0; i < int(math.Abs(float64(repeat.int))); i++ {
		do()
	}
}
