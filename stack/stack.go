package stack

import "github.com/wang502/go-collections/linkedlist"

type Stack struct{
    Elements *linkedlist.LinkedList
}

func NewStack() *Stack{
    ll := linkedlist.NewLinkedList()
    return &Stack{ll}
}

func (stack *Stack) Push(val int) {
    ll := stack.Elements
    ll.AddHead(val)
}

func (stack *Stack) Pop() int {
    size := stack.Elements.Size()
    if size == 0 {
      return -1
    }
    ll := stack.Elements
    return ll.RemoveHead()
}

func (stack *Stack) Peek() int {
    size := stack.Elements.Size()
    if size == 0{
      return -1
    }
    return stack.Elements.GetHead()
}

func (stack *Stack) Size() int {
    return stack.Elements.Size()
}
