//Data Oriented style.
package dod

import (

)

type Bot struct {
	Position [3]float64
	Velocity [3]float64
}

func update(b []*Bot, count int, time float64) {
	for i := 0; i < count; i++ {
		for j := 0; j < 3; j++ {
			b[i].Position[j] += b[i].Velocity[j] * time
		}
	}
}