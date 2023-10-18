package util

import (
	"fmt"
	"strings"
)

func NewRing[T any]() *Ring[T] {
	return &Ring[T]{
		Head: nil,
	}
}

type Ring[T any] struct {
	Head *RingItem[T]
}

func (r *Ring[T]) String() string {
	if r.Head == nil {
		return ""
	}
	builder := strings.Builder{}

	item := r.Head
	for {
		builder.WriteString(fmt.Sprintf("%v", item.Value) + ", ")

		item = item.Next
		if item == r.Head {
			break
		}
	}

	return builder.String()
}

func (r *Ring[T]) Prepend(item T) {
	newItem := &RingItem[T]{
		Value: item,
	}
	if r.Head == nil {
		newItem.Next = newItem
		newItem.Previous = newItem
		r.Head = newItem
		return
	}

	newItem.Previous = r.Head.Previous
	newItem.Next = r.Head

	if r.Head.Previous != r.Head {
		r.Head.Previous.Previous.Next = newItem
	}
	r.Head.Previous = newItem

	if r.Head.Next == r.Head {
		r.Head.Next = newItem
	}

	r.Head = newItem
}

func (r *Ring[T]) Append(item T) {
	newItem := &RingItem[T]{
		Value: item,
	}

	if r.Head == nil {
		newItem.Next = newItem
		newItem.Previous = newItem
		r.Head = newItem
		return
	}

	oldTail := r.Head.Previous

	newItem.Next = r.Head
	newItem.Previous = oldTail
	oldTail.Next = newItem
	r.Head.Previous = newItem

	if r.Head.Next == r.Head {
		r.Head.Next = newItem
	}
}

type RingItem[T any] struct {
	Value    T
	Next     *RingItem[T]
	Previous *RingItem[T]
}
