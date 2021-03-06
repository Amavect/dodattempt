//Procedural inline style.
package inl

import (
	"testing"
	"fmt"
)

func TestDod(t *testing.T) {
	var count int = 1024
	var loops int = 100000
	var dt float64 = 1.0
	
	var bots []*Bot = make([]*Bot, count)
	for i := 0; i < count; i++ {
		bots[i] = &Bot{ [3]float64{0.0, 0.0, 0.0} , [3]float64{1.0,2.0,3.0} ,}
	}
	
	for i := 0; i < loops; i++ {
		for j := 0; j < count; j++ {
			for k := 0; k < 3; k++ {
				bots[j].Position[k] += bots[j].Velocity[k] * dt
			}
		}
	}
	fmt.Println(bots[0])
}