//Procedural style.
package pro

import (

)

type Bot struct {
	Position [3]float64
	Velocity [3]float64
}

func update(b *Bot, time float64) {
	for i := 0; i < 3; i++ {
		b.Position[i] += b.Velocity[i] * time
	}
}