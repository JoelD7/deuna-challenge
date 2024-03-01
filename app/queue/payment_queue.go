package queue

import (
	"github.com/JoelD7/deuna-challenge/app/models"
)

var (
	q *queue
)

func init() {
	q = &queue{}
}

type node struct {
	data *models.Payment
	next *node
}

type queue struct {
	head *node
	tail *node
	size int
}

func Add(item *models.Payment) {
	newNode := &node{data: item, next: nil}
	if q.tail == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
	q.size++
}

func RemoveFirst() *models.Payment {
	if q.head == nil {
		return nil
	}

	item := q.head.data
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}

	q.size--
	return item
}

// RemoveForMerchant removes the first payment in the queue for the given merchantID
func RemoveForMerchant(merchantAccountID string) *models.Payment {
	if q.head == nil {
		return nil
	}

	var prev *node
	curr := q.head
	for curr != nil {
		if *curr.data.MerchantAccountID == merchantAccountID {
			if prev == nil {
				return RemoveFirst()
			}

			prev.next = curr.next
			q.size--
			return curr.data
		}

		prev = curr
		curr = curr.next
	}
	return nil
}

func IsEmpty() bool {
	return q.size == 0
}

func Size() int {
	return q.size
}
