package main

// Request from modeling system
type Request struct {
	requests  int
	blocks    int
	isRequest bool
	blocked   bool
}

func (request *Request) generate() {
	if !request.blocked {
		request.isRequest = !request.isRequest
	}
}

func (request *Request) hasRequest() bool {
	return request.isRequest
}

func (request *Request) blocking() {
	request.blocks++
	request.blocked = true
}

func (request *Request) isBlocked() bool {
	return request.blocked
}

func (request *Request) unblocking() {
	request.blocked = false
}

func (request *Request) getBlocks() int {
	return request.blocks
}

func (request *Request) getRequest() int {
	if request.blocked {
		return 0
	} else if request.isRequest {
		return 1
	}
	return 2
}
