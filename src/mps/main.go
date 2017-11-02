package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const t1, t2, t3, t4 = 48.0 * 60, 6.0 * 60, 0.5 * 60, 1.0 * 60
	const iterations = 10000000
	var p [4]float64

	L := [4]float64{1 / t1, 1 / t3, 1 / (t2 - t3 - t4), 1 / t4}
	state := State{maxPosition: 3, p: L}

	for i := 0; i < iterations; i++ {
		state.generate()
		p[state.position]++
	}

	for i := 0; i < 4; i++ {
		fmt.Printf("P%d: %f\n", i, p[i]/iterations)
	}
}
