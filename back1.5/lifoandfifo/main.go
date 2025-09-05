package main

import "fmt"

type Stack struct {
    items []interface{}
}

func (s *Stack) Push(item interface{}) {
    s.items = append(s.items, item)
}

func (s *Stack) Pop() interface{} {
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item
}

type Queue struct {
    items []interface{}
}

func (q *Queue) Enqueue(item interface{}) {
    q.items = append(q.items, item)
}

func (q *Queue) Dequeue() interface{} {
    item := q.items[0]
    q.items = q.items[1:]
    return item
}

func main() {
    stack := &Stack{}
    stack.Push(1)
    stack.Push(2)
    stack.Push(3)
    fmt.Println(stack.Pop(), stack.Pop(), stack.Pop())

    queue := &Queue{}
    queue.Enqueue(1)
    queue.Enqueue(2)
    queue.Enqueue(3)
    fmt.Println(queue.Dequeue(), queue.Dequeue(), queue.Dequeue())
}