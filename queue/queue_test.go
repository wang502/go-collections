package queue

import (
  "testing"
)

const iterations = 10

func testEnqueueDequeue(t *testing.T){
    queue := NewQueue()
    for i:= 0; i<iterations; i++{
        queue.Enqueue(i)
    }

    for i:=0; i<iterations; i++{
      if e:=queue.Dequeue(); e!=i {
        t.Errorf("Queue Dequeue() expected to return %d, but returned %d", i, e)
      }
    }
}

func testSize(t *testing.T){
    queue := NewQueue()
    for i:= 0; i<iterations; i++{
      queue.Enqueue(i)
    }
    if queue.Size() != iterations{
      t.Errorf("Queue Size() expected to return %d, but returned %d", iterations, queue.Size())
    }
}

func testTop(t *testing.T){
    queue := NewQueue()
    for i:=0; i<iterations; i++{
      queue.Enqueue(i)
    }

    for i:=0; i<iterations; i++{
      if top:=queue.Top(); top!=i {
        t.Errorf("Queue Top() expected to return %d, but returned %d", i, top)
      }
      queue.Dequeue()
    }
}
