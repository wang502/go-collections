package stack

import "testing"

const iterations = 100

func testPushPop(t *testing.T){
  s := NewStack()
  for i:=0; i<iterations; i++{
    s.Push(i)
  }
  var top interface{}
  for j:=iterations-1; j>-1; j--{
    if s, top = s.Pop(); top!=j{
      t.Errorf("Stack Pop() error")
    }
  }
}
