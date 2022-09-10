package queue

import "testing"

func TestQueue(t *testing.T) {
	item := 1

	q := New[int]()
	q.Enqueue(&item)

	got := q.Dequeue()
	if *got != 1 {
		t.Fail()
	}
}
