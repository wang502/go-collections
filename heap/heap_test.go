package heap

import  (
    "testing"
    "strconv"
    "strings"
)

const iterations = 10

type TestItem int

func (t_item TestItem) Compare(other Item) int{
    o_item := other.(TestItem)
    if t_item < o_item{
        return -1
    } else if (t_item > o_item){
        return 1
    } else{
      return 0
    }
}

func (t_item TestItem) String() string {
    return strconv.Itoa(int(t_item))
}

func testInsert(t *testing.T){
    heap := NewHeap()
    for i:=iterations-1; i>-1; i--{
      heap.Insert(TestItem(i))
    }
    if (heap.Size()) != iterations{
      t.Errorf("Heap Insert() for 10 times, Size() expected to return %d, but returned %d", iterations, heap.Size())
    }
}

func testExtractMin(t *testing.T){
    heap := NewHeap()
    for i:=iterations-1; i>-1; i--{
      heap.Insert(TestItem(i))
    }

    cur_size := heap.Size()
    for i:=0; i<iterations; i++{
      if cur_size != heap.Size(){
        t.Errorf("Heap ExtractMin() expected to change size, but not")
      }

      if min:=heap.ExtractMin(); min.Compare(TestItem(i)) != 0{
        t.Errorf("Heap ExtractMin() expected to return %d, but returned %d", i, min)
      }
      cur_size -= 1;
    }
}

func testItems(t *testing.T){
  heap := NewHeap()
  for i:=iterations-1; i>-1; i--{
    heap.Insert(TestItem(i))
  }
  repr := ""
  for heap.Size()!= 0 {
      repr += heap.ExtractMin().String() + " "
  }

  res := ""
  for i:=0; i<iterations; i++{
    res += strconv.Itoa(i) + " "
  }
  if strings.Compare(res, repr) != 0{
    t.Errorf("Heap representation not correct")
  }
}
