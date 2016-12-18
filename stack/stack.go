package stack

type stack []interface{}

func NewStack() stack{
  return make(stack, 0)
}

func (s stack) Push(i interface{}) (stack){
  return append(s, i)
}

func (s stack) Pop() (stack, interface{}){
  l := len(s)
  return s[:l-1], s[l-1]
}
