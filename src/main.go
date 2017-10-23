package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const p1, p2 = 1 - 0.75, 1 - 0.7
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
				}
			}
		} else {
			calculate(&channel1, &channel2, &queue, &request)
		}
	}

	fmt.Println(channel1.processed)
	fmt.Println(channel2.processed)
	fmt.Println(request.getDiscards())
	processed := channel2.getProcessed() + channel1.getProcessed()
	fmt.Println(float64(processed) / float64(iterations))
	fmt.Println(0.5 * (1 - float64(request.getDiscards())/float64(iterations/2)))
}

func calculate(channel1 *Channel, channel2 *Channel, queue *Queue, request *Request) {
	if channel1.isEmpty() && !queue.isEmpty() {
		channel1.holding()
		queue.removeItem()
		if channel2.isEmpty() && !queue.isEmpty() {
			channel2.holding()
			queue.removeItem()
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
		}
	}
}
