package tree

import "fmt"

type Tree struct{
  Left *Tree
  Value int
  Right *Tree
}

func insert(tree *Tree, v int) *Tree{
  if tree == nil{
    return &Tree{nil, v, nil}
  }
  if tree.Value < v{
    tree.Right = insert(tree.Right, v)
  }else{
    tree.Left = insert(tree.Left, v)
  }
  return tree
}

func (tree *Tree) walk(){
  if tree == nil{
    return
  }
  tree.Left.walk()
  fmt.Println(tree.Value)
  tree.Left.walk()
}

func (tree *Tree) String() string{
  if tree == nil{
    return "()"
  }
  res := ""
  if tree.Left != nil{
    res += tree.Left.String() + " "
  }
  res += fmt.Sprint(tree.Value)
  if tree.Right != nil{
    res += " " + tree.Right.String()
  }
  return "(" + res + ")"
}
/*
func main(){
  var tree *Tree
  tree = insert(tree, 5)
  tree = insert(tree, 3)
  tree = insert(tree, 10)

  fmt.Println(tree.String())
  tree.walk()

  var nil_tree *Tree = nil
  nil_tree.String()

  var arr []int
  fmt.Println(len(arr))
}
*/
