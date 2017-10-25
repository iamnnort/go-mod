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

	request := Request{}
	queue := Queue{maxItems: 2}
	channel1 := Channel{p: p1}
	channel2 := Channel{p: p2}

	for i := 0; i < iterations; i++ {
		request.generate()
		if request.hasRequest() {
			if !queue.isFull() {
				queue.addItem()
				calculate(&channel1, &channel2, &queue, &request)
			} else {
				calculate(&channel1, &channel2, &queue, &request)
				if queue.isFull() {
					request.discarding()
				} else {
					queue.addItem()
				}
			}
		} else {
			calculate(&channel1, &channel2, &queue, &request)
		}
	}

	statistics(request, queue, channel1, channel2, iterations)
}

func calculate(channel1 *Channel, channel2 *Channel, queue *Queue, request *Request) {
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
						channel2.processing()
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
						channel2.processing()
						queue.removeItem()
					}
				}
			} else {
				channel1.processing()
				queue.removeItem()
				if !queue.isEmpty() {
					if channel2.isEmpty() {
						channel2.holding()
						queue.removeItem()
					} else {
						channel2.generate()
						if !channel2.isHold() {
							channel2.processing()
							queue.removeItem()
						}
					}
				}
			}
		}
	}
}

func statistics(request Request, queue Queue, channel1 Channel, channel2 Channel, iterations int) {
	pBlock := float64(request.getDiscards()) / float64(iterations/2)
	A := 0.5 * (1 - pBlock)
	A1 := float64(channel1.getProcessed()+channel2.getProcessed()) / float64(iterations/2)
	L := float64(queue.getSummary()) / float64(iterations/2)
	W := float64(L) / float64(A)

	fmt.Println("A:", A)
	fmt.Println("A1:", A1)
	fmt.Println("L:", L)
	fmt.Println("W:", W)
}
