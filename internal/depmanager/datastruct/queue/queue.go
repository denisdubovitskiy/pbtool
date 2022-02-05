package queue

import "container/list"

type Queue struct {
	q *list.List
}

func New() *Queue {
	return &Queue{q: list.New()}
}

func (q *Queue) Enqueue(e string) {
	q.q.PushBack(e)
}

func (q *Queue) Dequeue() string {
	val := q.q.Front()
	defer q.q.Remove(val)
	return val.Value.(string)
}

func (q *Queue) Len() int {
	return q.q.Len()
}
