package main

import "math/rand"

//Channel from modeling system
type Channel struct {
	processed int
	discards  int
	hold      bool
	value     int
	p         float64
}

func (channel *Channel) getValue() int {
	return channel.value
}

func (channel *Channel) isEmpty() bool {
	return channel.value == 0
}

func (channel *Channel) isFull() bool {
	return channel.value == 1
}

func (channel *Channel) holding() {
	channel.value = 1
}

func (channel *Channel) processing() {
	channel.value = 0
	channel.processed++
}

func (channel *Channel) isHold() bool {
	return channel.hold
}

func (channel *Channel) getProcessed() int {
	return channel.processed
}

func (channel *Channel) generate() {
	channel.hold = rand.Float64() <= channel.p
}
