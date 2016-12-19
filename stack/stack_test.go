package stack

import "testing"

const iterations = 10

func testPushPop(t *testing.T){
    s := NewStack()
    for i:=0; i<iterations; i++{
      s.Push(i)
    }
    var top interface{}
    for j:=iterations-1; j>-1; j--{
      if top = s.Pop(); top!=j{
        t.Errorf("Stack Pop() error")
      }
    }
}

func testPeek(t *testing.T){
    s := NewStack()
    for i:=0; i<iterations; i++{
      s.Push(i)
      if top := s.Peek(); top != i{
        t.Errorf("Stack Top() expected to return %d, but returns %d", i, top)
      }
    }
}

func testSize(t *testing.T){
    s := NewStack()
    for i:=0; i<iterations; i++{
      s.Push(i)
    }
    if s.Size() != iterations{
      t.Errorf("Stack Size() expected to return %d, but returns %d", iterations, s.Size())
    }
}
