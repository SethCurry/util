package util_test

import (
	"testing"

	"github.com/SethCurry/util"
	"github.com/stretchr/testify/assert"
)

func Test_LinkedList_Slice(t *testing.T) {
	l := util.NewLinkedList[int]()
	l.Prepend(1)
	l.Prepend(2)
	l.Prepend(3)

	asSlice := l.Slice()
	assert.Equal(t, []int{3, 2, 1}, asSlice)

	l.Remove(l.Head.Next)
	assert.Equal(t, []int{3, 1}, l.Slice())

	l.Append(4)
	assert.Equal(t, []int{3, 1, 4}, l.Slice())
}
