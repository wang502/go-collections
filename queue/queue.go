package queue

import (
    "github.com/wang502/go-collections/linkedlist"
)

type Queue struct{
    Elements *linkedlist.LinkedList
}

func NewQueue() *Queue{
    ll := linkedlist.NewLinkedList()
    return &Queue{ll}
}

func (queue *Queue) Enqueue(val interface{}) {
    queue.Elements.AddTail(val)
}

func (queue *Queue) Dequeue() interface{} {
    ll := queue.Elements
    if ll.Tail == nil{
        return nil
    }
    return ll.RemoveHead()
}

func (queue *Queue) Size() int {
    return queue.Elements.Size()
}

func (queue *Queue) Top() interface{} {
    ll := queue.Elements
    if ll.Tail == nil{
      return nil
    }
    res := ll.Head.Val
    return res
}

func (queue *Queue) IsEmpty() bool {
  return queue.Elements.Size() == 0
}
