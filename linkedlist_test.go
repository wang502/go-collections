package data_structure

import "testing"
import "github.com/wang502/go-tour/data_structure"

const iterations = 100

func TestAdd(t *testing.T){
  ll := data_structure.NewLinkedList()
  for i:=0; i<iterations; i++{
    ll.Add(i)
  }
  if ll.Size() != iterations{
    t.Errorf("Linked List size expected to be %d, but is %d", iterations, ll.Size())
  }
}

func TestSearch(t *testing.T){
  ll := data_structure.NewLinkedList()
  for i:=0; i<iterations; i++{
    ll.Add(i)
  }
  for j:=0; j<iterations; j++{
    if !ll.Search(j){
      t.Errorf("Linked list Search expected to be True, but is False")
    }
  }
}

func TestReverse(t *testing.T){
  ll := data_structure.NewLinkedList()
  for i:= 0; i<iterations; i++{
    ll.Add(i)
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
