package data_structure

import "testing"
import "github.com/wang502/go-tour/data_structure"


func testPushPop(t *testing.T){
  s := data_structure.NewStack()
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
