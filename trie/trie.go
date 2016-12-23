package trie

type TrieNode struct{
    children map[rune]*TrieNode
    is_word bool
    word string
}

type Trie struct{
    root *TrieNode
}

func NewTrie() *Trie {
    rootNode := &TrieNode{
              children: make(map[rune]*TrieNode),
              is_word: false,
              word: "",
              }
    return &Trie{
              root: rootNode,
           }
}

func (trie *Trie) Add(word string) bool{
    curNode := trie.root
    for _, r := range word{
        _, ok := curNode.children[r]
        if !ok{
            curNode.children[r] = &TrieNode{make(map[rune]*TrieNode), false, ""}
        }
        curNode = curNode.children[r]
    }
    curNode.is_word = true
    curNode.word = word
    return true
}

func (trie *Trie) Search(word string) bool{
    curNode := trie.root
    for _, r := range word{
        _, ok := curNode.children[r]
        if !ok{
            return false
        }
        curNode = curNode.children[r]
    }
    return curNode.is_word
}
