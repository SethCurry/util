package util

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

type LinkedList[T any] struct {
	Head *LinkedListItem[T]
	Tail *LinkedListItem[T]
	Len  int
}

func (l *LinkedList[T]) Prepend(value T) {
	oldHead := l.Head
	newItem := &LinkedListItem[T]{
		Next:  oldHead,
		Value: value,
	}
	l.Head = newItem
	if oldHead != nil {
		oldHead.Previous = newItem
	}

	if l.Tail == nil {
		l.Tail = newItem
	}
	l.Len++
}

func (l *LinkedList[T]) Append(value T) {
	oldTail := l.Tail
	newItem := &LinkedListItem[T]{
		Previous: oldTail,
		Value:    value,
	}
	l.Tail = newItem
	if oldTail != nil {
		oldTail.Next = newItem
	}

	if l.Head == nil {
		l.Head = newItem
	}
	l.Len++
}

func (l *LinkedList[T]) Remove(item *LinkedListItem[T]) {
	item.Previous.Next = item.Next
	item.Next.Previous = item.Previous
	l.Len--
}

func (l *LinkedList[T]) Slice() []T {
	ret := make([]T, l.Len)

	item := l.Head

	for i := 0; i < l.Len; i++ {
		ret[i] = item.Value
		item = item.Next
	}

	return ret
}

type LinkedListItem[T any] struct {
	Next     *LinkedListItem[T]
	Previous *LinkedListItem[T]
	Value    T
}
