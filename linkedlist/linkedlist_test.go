package linkedlist

import("testing"
       "strconv"
       "strings"
      )

const iterations = 10

func TestAddTail(t *testing.T){
  ll := NewLinkedList()
  for i:=0; i<iterations; i++{
    ll.AddTail(i)
  }

  s := ll.String()
  res := ""
  for i:=0; i<iterations; i++{
    res += strconv.Itoa(i) + "->"
  }
  res += "None\n"
  if strings.Compare(s, res) != 0{
    t.Errorf("Linked List expected to contain every added elements, but not")
  }

  if ll.Size() != iterations{
    t.Errorf("Linked List size expected to be %d, but is %d", iterations, ll.Size())
  }
}

func TestAddHead(t *testing.T){
  ll := NewLinkedList()
  for i:=0; i<iterations; i++{
    ll.AddHead(i)
  }

  s := ll.String()
  res := ""
  for i:=iterations-1; i>-1; i--{
    res += strconv.Itoa(i) + "->"
  }
  res += "None\n"
  if strings.Compare(s, res) != 0{
    t.Errorf("Linked List expected to contain every added elements, but not")
  }

  if ll.Size() != iterations{
    t.Errorf("Linked List size expected to be %d, but is %d", iterations, ll.Size())
  }
}

func TestSearch(t *testing.T){
  ll := NewLinkedList()
  for i:=0; i<iterations; i++{
    ll.AddTail(i)
  }
  for j:=0; j<iterations; j++{
    if !ll.Search(j){
      t.Errorf("Linked list Search expected to be True, but is False")
    }
  }
}

func TestReverse(t *testing.T){
  ll := NewLinkedList()
  for i:= 0; i<iterations; i++{
    ll.AddTail(i)
  }
  ll.Reverse()
  cur := ll.Head
  for j:=iterations-1; j>-1; j--{
    if cur.Val != j{
      t.Errorf("Linked List Reverse Error")
    }
    cur = cur.Next
  }
}
