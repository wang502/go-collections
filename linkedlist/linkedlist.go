package linkedlist

import("fmt"
       "errors"
       "strconv"
)

type Node struct{
  Val int
  Prev *Node
  Next *Node
}

type LinkedList struct{
  Head *Node
  Tail *Node
  Length int
}

func NewLinkedList() *LinkedList{
  return &LinkedList{nil, nil, 0}
}

func (list *LinkedList) AddTail(val int) *LinkedList{
  newNode := &Node{val, list.Tail, nil}
  if list.Tail != nil{
    list.Tail.Next = newNode
  }else{
    list.Head = newNode
  }
  list.Tail = newNode
  list.Length += 1
  return list
}

func (list *LinkedList) AddHead(val int) *LinkedList{
  newNode := &Node{val, nil, list.Head}
  if list.Head == nil{
    list.Tail = newNode
  } else{
    list.Head.Prev = newNode
  }
  list.Head = newNode
  list.Length += 1
  return list
}

func (list *LinkedList) AddNode(node *Node) *LinkedList{
  if list.Tail != nil{
    node.Prev = list.Tail
    list.Tail.Next = node
  }else{
    list.Head = node
  }
  list.Tail = node
  list.Length += 1
  return list
}

func (list *LinkedList) Delete(node *Node) (*LinkedList, error){
  if list.Size() == 0{
    return list, errors.New("Linked list is empty")
  }
  if list.Head == list.Tail && list.Head == node{
    list.Head = nil
    list.Tail = nil
  }else if list.Head == node{
    list.Head = list.Head.Next
    list.Head.Next = nil
  } else if list.Tail == node{
    list.Tail = list.Tail.Prev
    list.Tail.Next = nil
  }else{
    node.Next.Prev = node.Prev
    node.Prev.Next = node.Next
  }
  list.Length -= 1
  return list, nil
}

func (list *LinkedList) Search(val int) bool{
  cur := list.Head
  for cur != nil{
    if cur.Val == val{
      return true
    }
    cur = cur.Next
  }
  return false
}

func (list *LinkedList) Reverse() *LinkedList{
  prev_head := list.Head
  cur := list.Head
  var reversed *Node
  for cur != nil{
    tmp := cur.Next
    cur.Next = reversed
    reversed = cur
    cur = tmp
  }
  list.Head = reversed
  list.Tail = prev_head
  return list
}

func (list *LinkedList) Size() int{
  return list.Length
}

func (list *LinkedList) String() string{
  res := ""
  if list == nil{
    fmt.Println(res)
  }
  cur := list.Head
  for cur != nil{
    res += strconv.Itoa(cur.Val) + "->"
    cur = cur.Next
  }
  res += "None\n"
  return res
}
