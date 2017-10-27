package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const p1, p2 = 0.75, 0.7
	const iterations = 10000000
	var p [3][3][2][2]int

	request := Request{}
	queue := Queue{maxItems: 2}
	channel1 := Channel{p: p1}
	channel2 := Channel{p: p2}

	for i := 0; i < iterations; i++ {
		request.generate()
		// fmt.Println(request.getRequest(), queue.getValue(), channel1.getValue(), channel2.getValue())
		p[request.getRequest()][queue.getValue()][channel1.getValue()][channel2.getValue()]++
		if request.hasRequest() {
			if !queue.isFull() {
				queue.addItem()
				calculate(&channel1, &channel2, &queue, &request)
			} else {
				calculate(&channel1, &channel2, &queue, &request)
				if queue.isFull() {
					request.blocking()
				} else {
					request.unblocking()
					queue.addItem()
				}
			}
		} else {
			calculate(&channel1, &channel2, &queue, &request)
		}
	}

	A := 0.5 * (1 - float64(p[0][2][1][1])/float64(iterations))
	L := float64(queue.getSummary()) / float64(iterations/2)
	W := float64(L) / float64(A)

	fmt.Println("A:", A)
	fmt.Println("L:", L)
	fmt.Println("W:", W)
}

func calculate(channel1 *Channel, channel2 *Channel, queue *Queue, request *Request) {
	if !request.isBlocked() {
		if !queue.isEmpty() {
			if channel1.isEmpty() {
				channel1.holding()
				queue.removeItem()
				if !queue.isEmpty() {
					if channel2.isEmpty() {
						channel2.holding()
						queue.removeItem()
					} else {
						channel2.generate()
						if !channel2.isHold() {
							queue.removeItem()
						}
					}
				}
			} else {
				channel1.generate()
				if channel1.isHold() {
					if channel2.isEmpty() {
						channel2.holding()
						queue.removeItem()
					} else {
						channel2.generate()
						if !channel2.isHold() {
							queue.removeItem()
						}
					}
				} else {
					queue.removeItem()
					if !queue.isEmpty() {
						if channel2.isEmpty() {
							channel2.holding()
							queue.removeItem()
						} else {
							channel2.generate()
							if !channel2.isHold() {
								queue.removeItem()
							}
						}
					}
				}
			}
		} else {
			if channel1.isFull() {
				channel1.generate()
				if !channel1.isHold() {
					channel1.processing()
				}
			}
			if channel2.isFull() {
				channel2.generate()
				if !channel2.isHold() {
					channel2.processing()
				}
			}
		}
	} else {
		channel1.generate()
		channel2.generate()
		if !channel1.isHold() && !channel2.isHold() {
			queue.removeItem()
			queue.removeItem()
			request.unblocking()
		}
	}
}
