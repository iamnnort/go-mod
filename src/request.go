package main

import "math/rand"

// Request from modeling system
type Request struct {
	requests   int
	p          int
	discards   int
	hasRequest bool
}

func (request *Request) generate() {
	if rand.Int() > request.p {
		request.p++
		request.hasRequest = true
	} else {
		request.hasRequest = false
	}
}
