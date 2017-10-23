package main

// Queue from modeling system
type Queue struct {
	maxItems int
	items    int
}

func (queue *Queue) addItem() {
	if queue.items < queue.maxItems {
		queue.items++
	}
}

func (queue *Queue) removeItem() {
	if queue.items > 0 {
		queue.items--
	}
}

func (queue *Queue) isFull() bool {
	return queue.items == queue.maxItems
}

func (queue *Queue) isEmpty() bool {
	return queue.items == 0
}

func (queue *Queue) getValue() int {
	return queue.items
}
