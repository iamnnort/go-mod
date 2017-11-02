package main

import "math/rand"

// State is mps system state
type State struct {
	position    int
	maxPosition int
	processed   bool
	p           [4]float64
}

func (state *State) generate() {
	state.processed = rand.Float64() < state.p[state.position]
	if state.isProcessed() {
		state.nextPosition()
	}
}

func (state *State) isProcessed() bool {
	return state.processed
}

func (state *State) nextPosition() {
	if state.position < state.maxPosition {
		state.position++
	} else {
		state.position = 0
	}
}
