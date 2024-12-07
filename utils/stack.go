package utils

type Stack[T any] struct {
	items []T
}

func NewStack[T any](items []T) *Stack[T] {
	return &Stack[T]{
		items: items,
	}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		panic("stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack[T]) Peek() T {
	if len(s.items) == 0 {
		panic("stack is empty")
	}
	return s.items[len(s.items)-1]
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

func (s *Stack[T]) Clone() *Stack[T] {
	items := make([]T, len(s.items))
	copy(items, s.items)
	return &Stack[T]{
		items: items,
	}
}
