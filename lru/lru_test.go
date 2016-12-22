package lru

import(
    "testing"
)

const iterations = 10

func TestNewLRU(t *testing.T)  {
    lru1 := NewLRU(10)
    if lru1 == nil{
        t.Errorf("NewLRU() expected to return pointer to LRU struct, but returned nil instead")
    }

    lru2 := NewLRU(-1)
    if lru2 != nil{
      t.Errorf("NewLRU() expected to return nil, but returned pointer to LRU struct instead")
    }
}

func TestAdd(t *testing.T){
    lru := NewLRU(iterations-2)
    for i:=0; i<iterations; i++{
        lru.Add(i, i)
    }

    for i:=0; i<2; i++{
        if lru.Contains(i){
            t.Errorf("LRU expected to have exivted %d", i)
        }
    }

    // test duplicate key
    lru.Add(iterations-1, 100)
    v := lru.Get(iterations-1)
    if v != 100 {
        t.Errorf("LRU Add() did not update value of existing key %d", iterations-1)
    }
}

func TestGet(t *testing.T){
    lru := NewLRU(iterations)
    for i:=0; i<iterations; i++{
        lru.Add(i, i)
    }

    for i:=0; i<iterations; i++{
        val := lru.Get(i)
        if val != lru.Peek().Value{
            t.Errorf("LRU Get() expected to place the item to head, but did not")
        }
    }
}

func TestKeys(t *testing.T){
    lru := NewLRU(iterations)
    for i:=0; i<iterations; i++{
        lru.Add(i, i)
    }

    keys := lru.Keys()

    i:=0
    for j:=iterations-1; j>-1; j--{
        if keys[j] != i{
            t.Errorf("LRU Keys() order not correct")
        }
        i += 1
    }
}
