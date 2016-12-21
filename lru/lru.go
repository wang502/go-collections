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
    return &LRU{ll, make(map[interface{}]*linkedlist.Node), size}
}

func (lru *LRU) Add(key interface{}, value interface{}) bool{
    if ent, ok := lru.Map[key]; ok{
        lru.Items.Detach(ent)
        lru.Items.PushFront(ent)
        ent.Val = value
        return false
    }

    newItem := &linkedlist.Node{Item{key, value}, nil, nil}
    lru.Items.PushFront(newItem)
    lru.Map[key] = newItem

    if lru.Items.Size() > lru.MAX_SIZE{
        lru.RemoveOldest()
    }
    return true
}

func (lru *LRU) RemoveOldest(){
    oldest := lru.Items.PopTail()
    if oldest == nil{
        return
    }
    key := oldest.Val.(Item).Key
    delete(lru.Map, key)
    return
}

func (lru *LRU) Contains(key interface{}) bool {
    _, ok := lru.Map[key]
    return ok
}
