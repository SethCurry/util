package util_test

import (
	"testing"

	"github.com/SethCurry/util"
	"github.com/stretchr/testify/assert"
)

func Test_Ring_Prepend(t *testing.T) {
	l := util.NewRing[int]()
	assert.Nil(t, l.Head)

	l.Prepend(1)
	assert.Equal(t, 1, l.Head.Value)

	l.Prepend(2)
	assert.Equal(t, 2, l.Head.Value)
	assert.Equal(t, 1, l.Head.Next.Value)
	assert.Equal(t, 1, l.Head.Previous.Value)
	assert.Equal(t, 2, l.Head.Next.Next.Value)
	assert.Equal(t, 2, l.Head.Previous.Previous.Value)

	l.Prepend(3)
	assert.Equal(t, 3, l.Head.Value)
	assert.Equal(t, 1, l.Head.Previous.Value)
	assert.Equal(t, 2, l.Head.Next.Value)
}

func Test_Ring_Append(t *testing.T) {
	l := util.NewRing[int]()
	assert.Nil(t, l.Head)

	l.Append(1)
	// 1
	assert.Equal(t, 1, l.Head.Value)

	l.Append(2)
	// 1, 2
	assert.Equal(t, 1, l.Head.Value)
	assert.Equal(t, 2, l.Head.Previous.Value)
	assert.Equal(t, 1, l.Head.Previous.Previous.Value)
	assert.Equal(t, 1, l.Head.Next.Next.Value)

	l.Append(3)
	// 1, 2, 3
	assert.Equal(t, 3, l.Head.Previous.Value)
	assert.Equal(t, 2, l.Head.Previous.Previous.Value)
	assert.Equal(t, 3, l.Head.Next.Next.Value)
}
