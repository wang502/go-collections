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
}
