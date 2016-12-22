package main

import(
    "fmt"
    "github.com/wang502/go-collections/lru"
)

func main(){
    lru := lru.NewLRU(3)
    lru.Add("1", 1)
    lru.Add("2", 2)
    lru.Add("3", 3)

    peek1 := lru.Peek()
    fmt.Printf("Head value is %d after inserting {'3':3}\n", peek1.Value)

    fmt.Printf("Value of '1' is %d \n", lru.Get("1"))

    peek2 := lru.Peek()
    fmt.Printf("Head value is %d after geting {'1':1}\n", peek2.Value)

    lru.Add("4", 4)
    peek3 := lru.Peek()
    fmt.Printf("Head value is %d after inserting {'4':4}\n", peek3.Value)

    fmt.Printf("LRU current size is %d\n", lru.Size())

    lru.Add("4", 40)
    fmt.Printf("Value of Key '4' is updated to be %d\n", lru.Get("4"))
}
