//Object Oriented style.
package oop

import (

)

type Bot struct {
	Position [3]float64
	Velocity [3]float64
}

func (b *Bot) update(time float64) {
	for i := 0; i < 3; i++ {
		b.Position[i] += b.Velocity[i] * time
	}
}