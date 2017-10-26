package main

// Request from modeling system
type Request struct {
	requests  int
	blocks    int
	isRequest bool
}

func (request *Request) generate() {
	request.isRequest = !request.isRequest
}

func (request *Request) hasRequest() bool {
	return request.isRequest
}

func (request *Request) blocking() {
	request.blocks++
}

func (request *Request) getBlocks() int {
	return request.blocks
}

func (request *Request) getRequest() int {
	if request.isRequest {
		return 1
	}
	return 2
}
