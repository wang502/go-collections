package trie

import (
    "testing"
)

func TestAddAndSeach(t *testing.T){
    trie := NewTrie()
    words1 := []string{"Merry", "Christmas", "Day"}
    for _, word := range words1 {
        trie.Add(word)
    }

    for _, word := range words1 {
        if !trie.Search(word){
            t.Errorf("Trie Search() expected to return true, but returned false")
        }
    }

    words2 := []string{"Happy", "New", "Year"}
    for _, word := range words2 {
        if trie.Search(word){
            t.Errorf("Trie Search() expected to return false, but returned true")
        }
    }
}
