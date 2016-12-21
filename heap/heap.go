package heap

type Item interface{
    Compare(other Item) int
    String() string
}

type Heap struct {
    Elements []Item
}

func NewHeap() *Heap {
    elements := make([]Item, 0)
    return &Heap{elements}
}

func (heap *Heap) Insert(val Item){
    heap.Elements = append(heap.Elements, val)
    heap.DecreaseKey(len(heap.Elements)-1)
}

func parent(i int) int{
    return (i-1)/2
}

func swap(elements []Item, i int, j int)  {
    tmp := elements[i]
    elements[i] = elements[j]
    elements[j] = tmp
}

func (heap *Heap) DecreaseKey(index int) {
    for index>=0 && parent(index) >= 0 && heap.Elements[index].Compare(heap.Elements[parent(index)]) < 0{
        swap(heap.Elements, index, parent(index))
        index = parent(index)
    }
}

func (heap *Heap) ExtractMin() Item{
    if len(heap.Elements) == 0{
      return nil
    }
    res := heap.Elements[0]
    size := len(heap.Elements)
    heap.Elements[0] = heap.Elements[size-1]
    heap.Elements = append(heap.Elements[:size-1])
    heap.MinHeapify(0)
    return res
}

func (heap *Heap) GetMin() Item {
    if len(heap.Elements) == 0{
        return nil
    }
    return heap.Elements[0]
}

func (heap *Heap) MinHeapify(index int) {
    elements := heap.Elements
    min_i := index
    left, right := (2*index)+1, (2*index)+2
    if left < len(elements) && elements[left].Compare(elements[min_i]) < 0{
        min_i = left
    }
    if right < len(elements) && elements[right].Compare(elements[min_i]) < 0{
        min_i = right
    }
    if min_i != index{
      tmp := elements[min_i]
      elements[min_i] = elements[index]
      elements[index] = tmp
      heap.MinHeapify(min_i)
    }
}

func (heap *Heap) BuildMinHeap(){
    elements := heap.Elements
    last_parent := len(elements)/2-1
    for i:=last_parent; i>-1; i--{
      heap.MinHeapify(i)
    }
}

func (heap *Heap) Size() int{
    return len(heap.Elements)
}
