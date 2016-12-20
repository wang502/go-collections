package queue

import (
    "github.com/wang502/go-collections/linkedlist"
)

type Queue struct{
    Elements *linkedlist.LinkedList
}

func NewQueue() *Queue {
    ll := linkedlist.NewLinkedList()
    return &Queue{ll}
}

func (queue *Queue) Enqueue(val int) {
    queue.Elements.AddTail(val)
}

func (queue *Queue) Dequeue() int {
    ll := queue.Elements
    if ll.Tail == nil{
        return -1
    }
    return ll.RemoveHead()
}

func (queue *Queue) Size() int {
    return queue.Elements.Size()
}

func (queue *Queue) Top() int {
    ll := queue.Elements
    if ll.Tail == nil{
      return -1
    }
    res := ll.Head.Val
    return res
}
