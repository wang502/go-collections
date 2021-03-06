package lru

import(
  "github.com/wang502/go-collections/linkedlist"
  "fmt"
)

type Item struct{
    Key interface{}
    Value interface{}
}

type LRU struct{
    Items *linkedlist.LinkedList
    Map map[interface{}]*linkedlist.Node
    MAX_SIZE int
}

func NewLRU(size int) *LRU{
    if size < 0{
        fmt.Println("invalid size")
        return nil
    }
    ll := linkedlist.NewLinkedList()
    return &LRU{
              Items: ll,
              Map: make(map[interface{}]*linkedlist.Node),
              MAX_SIZE: size,
            }
}

func (lru *LRU) Add(key interface{}, value interface{}) bool{
    if ent, ok := lru.Map[key]; ok{
        lru.Items.Detach(ent)
        lru.Items.PushFront(ent)
        ent.Val.(*Item).Value = value
        //fmt.Printf("v:%d", ent.Val.(*Item).Value)
        return false
    }

    newItem := &linkedlist.Node{&Item{key, value}, nil, nil}
    lru.Items.PushFront(newItem)
    lru.Map[key] = newItem

    if lru.Items.Size() > lru.MAX_SIZE{
        lru.RemoveOldest()
    }
    return true
}

func (lru *LRU) Get(key interface{}) interface{} {
    node, ok := lru.Map[key]
    var item *Item
    if !ok{
        return nil
    }
    item = node.Val.(*Item)
    lru.Items.Detach(node)
    lru.Items.PushFront(node)
    return item.Value
}

func (lru *LRU) Remove(key interface{}) bool {
    node, bool := lru.Map[key]
    if !bool{
        return false
    }
    lru.Items.Detach(node)
    delete(lru.Map, key)
    return true
}
func (lru *LRU) RemoveOldest(){
    oldest := lru.Items.PopTail()
    if oldest == nil{
        return
    }
    key := oldest.Val.(*Item).Key
    delete(lru.Map, key)
    return
}

func (lru *LRU) Contains(key interface{}) bool {
    _, ok := lru.Map[key]
    return ok
}

func (lru *LRU) Peek() Item{
    head := lru.Items.GetHead()
    p_item := head.Val.(*Item)
    return *p_item
}

func (lru *LRU) Keys() []interface{}{
    keys := make([]interface{}, lru.Size())
    i := 0
    for node := lru.Items.GetHead(); node != nil; node = node.Next {
        keys[i] = node.Val.(*Item).Value
        i += 1
    }
    return keys
}

func (lru *LRU) Size() int {
    return lru.Items.Size()
}
