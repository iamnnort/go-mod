package main

// Request from modeling system
type Request struct {
	requests  int
	discards  int
	isRequest bool
}

func (request *Request) generate() {
	request.isRequest = !request.isRequest
}

func (request *Request) hasRequest() bool {
	return request.isRequest
}

func (request *Request) discarding() {
	request.discards++
}

func (request *Request) getDiscards() int {
	return request.discards
}
