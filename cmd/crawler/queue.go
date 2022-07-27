package main

func newQueue() queue {
	return queue{el: make([]string, 0)}
}

type queue struct {
	el []string
}

func (q *queue) top() (string, bool) {
	if q.isEmpty() {
		return "", false
	}
	return q.el[0], true
}

func (q *queue) pop() (string, bool) {
	if q.isEmpty() {
		return "", false
	}
	val := q.el[0]
	q.el = q.el[1:]
	return val, true
}

func (q *queue) push(values ...string) {
	q.el = append(q.el, values...)
}

func (q *queue) isEmpty() bool {
	return len(q.el) == 0
}

func (q *queue) length() int {
	return len(q.el)
}

func (q *queue) contains(val string) bool {
	for _, el := range q.el {
		if el == val {
			return true
		}
	}
	return false
}
