package dsa_test

import (
	"slices"
	"testing"

	dsa "github.com/MilosSlavic/go-playground/pkg/linkedlist"
	"github.com/stretchr/testify/assert"
)

func TestNewLinkedListUsingSlice(t *testing.T) {
	data := make([]any, 3)
	data[0] = 1
	data[1] = 2
	data[2] = 3
	ll := dsa.NewLinkedListUsingSlice(&data)

	assert.Equal(t, 3, ll.GetCount())
	assert.Equal(t, []any{1, 2, 3}, slices.Collect(ll.All()))
}

func TestAdd(t *testing.T) {
	ll := dsa.NewLinkedList()
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)

	assert.Equal(t, 3, ll.GetCount())
}

func TestAddFirst(t *testing.T) {
	ll := dsa.NewLinkedList()
	ll.Add(100)
	ll.AddFirst(99)

	first := ll.GetFirst()
	assert.Equal(t, 99, first)
}

func TestDelete(t *testing.T) {
	ll := dsa.NewLinkedList()
	ll.Add(100)
	ll.Add(99)

	ll.Delete(100)

	first := ll.GetFirst()

	assert.Equal(t, 99, first)
	assert.Equal(t, 1, ll.GetCount())
}

func TestDeleteAt_Middle(t *testing.T) {
	ll := dsa.NewLinkedList()
	ll.Add(100)
	ll.Add(99)
	ll.Add(98)

	ll.DeleteAt(1)

	assert.Equal(t, 2, ll.GetCount())
	assert.Equal(t, []any{100, 98}, slices.Collect(ll.All()))
	assert.Equal(t, 100, ll.GetFirst())
	assert.Equal(t, 98, ll.GetLast())
}

func TestDeleteAt_First(t *testing.T) {
	ll := dsa.NewLinkedList()
	ll.Add(100)
	ll.Add(99)
	ll.Add(98)

	ll.DeleteAt(0)

	assert.Equal(t, 2, ll.GetCount())
	assert.Equal(t, []any{99, 98}, slices.Collect(ll.All()))
	assert.Equal(t, 99, ll.GetFirst())
	assert.Equal(t, 98, ll.GetLast())
}

func TestDeleteAt_Last(t *testing.T) {
	ll := dsa.NewLinkedList()
	ll.Add(100)
	ll.Add(99)
	ll.Add(98)

	ll.DeleteAt(2)

	assert.Equal(t, 2, ll.GetCount())
	assert.Equal(t, []any{100, 99}, slices.Collect(ll.All()))
	assert.Equal(t, 100, ll.GetFirst())
	assert.Equal(t, 99, ll.GetLast())
}

func TestDeleteAt_SingleElement(t *testing.T) {
	ll := dsa.NewLinkedList()
	ll.Add(100)

	ll.DeleteAt(0)

	assert.Equal(t, 0, ll.GetCount())
	assert.Empty(t, slices.Collect(ll.All()))

	ll.Add(99)
	assert.Equal(t, 99, ll.GetFirst())
	assert.Equal(t, 99, ll.GetLast())
}

func TestDeleteAt_OutOfBounds(t *testing.T) {
	ll := dsa.NewLinkedList()
	ll.Add(100)
	ll.Add(99)

	ll.DeleteAt(-1)
	ll.DeleteAt(2)

	assert.Equal(t, 2, ll.GetCount())
	assert.Equal(t, []any{100, 99}, slices.Collect(ll.All()))
}

func TestDeleteLast(t *testing.T) {
	ll := dsa.NewLinkedList()
	ll.Add(100)
	ll.Add(99)
	ll.Add(98)

	ll.DeleteLast()

	assert.Equal(t, 2, ll.GetCount())
	assert.Equal(t, []any{100, 99}, slices.Collect(ll.All()))
	assert.Equal(t, 99, ll.GetLast())

	ll.DeleteLast()
	ll.DeleteLast()

	assert.Equal(t, 0, ll.GetCount())
	assert.Empty(t, slices.Collect(ll.All()))
}

func TestGetFirst(t *testing.T) {
	ll := dsa.NewLinkedList()
	ll.Add(100)
	ll.Add(99)

	assert.Equal(t, 100, ll.GetFirst())
}

func TestGetLast(t *testing.T) {
	ll := dsa.NewLinkedList()
	ll.Add(100)
	ll.Add(99)

	assert.Equal(t, 99, ll.GetLast())
}

func TestGetAt(t *testing.T) {
	ll := dsa.NewLinkedList()
	ll.Add(100)
	ll.Add(99)
	ll.Add(98)

	assert.Equal(t, 99, ll.GetAt(1))
}

func TestExists(t *testing.T) {
	ll := dsa.NewLinkedList()
	ll.Add(100)
	ll.Add(99)

	assert.True(t, ll.Exists(99))
}

func TestClear(t *testing.T) {
	ll := dsa.NewLinkedList()
	ll.Add(100)
	ll.Add(99)

	ll.Clear()

	assert.False(t, ll.Exists(100))
	assert.Equal(t, 0, ll.GetCount())
}

func TestAll(t *testing.T) {
	ll := dsa.NewLinkedList()
	ll.Add(100)
	ll.Add(99)

	actual := slices.Collect(ll.All())

	assert.Equal(t, []any{100, 99}, actual)
}

func TestGetCount(t *testing.T) {
	ll := dsa.NewLinkedList()
	ll.Add(100)
	ll.Add(99)

	assert.Equal(t, 2, ll.GetCount())
}
