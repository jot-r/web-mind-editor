package common

type Stack[T any] struct {
	Push    func(T)
	Pop     func() T
	IsEmpty func() bool
	Peek    func() T
}

func NewStack[T any]() Stack[T] {
	slice := make([]T, 0)
	return Stack[T]{
		Push: func(i T) {
			slice = append(slice, i)
		},
		Pop: func() T {
			res := slice[len(slice)-1]
			slice = slice[:len(slice)-1]
			return res
		},
		Peek: func() T {
			return slice[len(slice)-1]
		},
		IsEmpty: func() bool {
			return len(slice) == 0
		},
	}
}
