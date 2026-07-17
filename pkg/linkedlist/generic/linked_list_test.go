package generic_test

import (
	"slices"
	"testing"

	generic "github.com/MilosSlavic/go-playground/pkg/linkedlist/generic"
	"github.com/stretchr/testify/assert"
)

func TestLinkedList_NewLinkedListUsingSlice(t *testing.T) {
	data := make([]any, 3)
	data[0] = 1
	data[1] = 2
	data[2] = 3
	ll := generic.NewLinkedListUsingSlice(&data)

	assert.Equal(t, 3, ll.GetCount())
	assert.Equal(t, []any{1, 2, 3}, slices.Collect(ll.All()))
}

func TestLinkedList_Add(t *testing.T) {
	ll := generic.NewLinkedList[int]()
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)

	assert.Equal(t, 3, ll.GetCount())
}

func TestLinkedList_AddFirst(t *testing.T) {
	ll := generic.NewLinkedList[int]()
	ll.Add(100)
	ll.AddFirst(99)

	first := ll.GetFirst()
	assert.Equal(t, 99, first)
}

func TestLinkedList_Delete(t *testing.T) {
	ll := generic.NewLinkedList[int64]()
	ll.Add(100)
	ll.Add(99)

	ll.Delete(100)

	first := ll.GetFirst()

	assert.Equal(t, int64(99), first)
	assert.Equal(t, 1, ll.GetCount())
}

func TestLinkedList_DeleteAt_Middle(t *testing.T) {
	ll := generic.NewLinkedList[int64]()
	ll.Add(100)
	ll.Add(99)
	ll.Add(98)

	ll.DeleteAt(1)

	assert.Equal(t, 2, ll.GetCount())
	assert.Equal(t, []int64{100, 98}, slices.Collect(ll.All()))
	assert.Equal(t, int64(100), ll.GetFirst())
	assert.Equal(t, int64(98), ll.GetLast())
}

func TestLinkedList_DeleteAt_First(t *testing.T) {
	ll := generic.NewLinkedList[int64]()
	ll.Add(100)
	ll.Add(99)
	ll.Add(98)

	ll.DeleteAt(0)

	assert.Equal(t, 2, ll.GetCount())
	assert.Equal(t, []int64{99, 98}, slices.Collect(ll.All()))
	assert.Equal(t, int64(99), ll.GetFirst())
	assert.Equal(t, int64(98), ll.GetLast())
}

func TestLinkedList_DeleteAt_Last(t *testing.T) {
	ll := generic.NewLinkedList[int64]()
	ll.Add(100)
	ll.Add(99)
	ll.Add(98)

	ll.DeleteAt(2)

	assert.Equal(t, 2, ll.GetCount())
	assert.Equal(t, []int64{100, 99}, slices.Collect(ll.All()))
	assert.Equal(t, int64(100), ll.GetFirst())
	assert.Equal(t, int64(99), ll.GetLast())
}

func TestLinkedList_DeleteAt_SingleElement(t *testing.T) {
	ll := generic.NewLinkedList[int64]()
	ll.Add(100)

	ll.DeleteAt(0)

	assert.Equal(t, 0, ll.GetCount())
	assert.Empty(t, slices.Collect(ll.All()))

	ll.Add(99)
	assert.Equal(t, int64(99), ll.GetFirst())
	assert.Equal(t, int64(99), ll.GetLast())
}

func TestLinkedList_DeleteAt_OutOfBounds(t *testing.T) {
	ll := generic.NewLinkedList[int64]()
	ll.Add(100)
	ll.Add(99)

	ll.DeleteAt(-1)
	ll.DeleteAt(2)

	assert.Equal(t, 2, ll.GetCount())
	assert.Equal(t, []int64{100, 99}, slices.Collect(ll.All()))
}

func TestLinkedList_DeleteLast(t *testing.T) {
	ll := generic.NewLinkedList[int64]()
	ll.Add(100)
	ll.Add(99)
	ll.Add(98)

	ll.DeleteLast()

	assert.Equal(t, 2, ll.GetCount())
	assert.Equal(t, []int64{100, 99}, slices.Collect(ll.All()))
	assert.Equal(t, int64(99), ll.GetLast())

	ll.DeleteLast()
	ll.DeleteLast()

	assert.Equal(t, 0, ll.GetCount())
	assert.Empty(t, slices.Collect(ll.All()))
}

func TestLinkedList_GetFirst(t *testing.T) {
	ll := generic.NewLinkedList[string]()
	ll.Add("100")
	ll.Add("99")

	assert.Equal(t, "100", ll.GetFirst())
}

func TestLinkedList_GetLast(t *testing.T) {
	ll := generic.NewLinkedList[string]()
	ll.Add("100")
	ll.Add("99")

	assert.Equal(t, "99", ll.GetLast())
}

func TestLinkedList_GetAt(t *testing.T) {
	ll := generic.NewLinkedList[string]()
	ll.Add("100")
	ll.Add("99")
	ll.Add("98")

	assert.Equal(t, "99", ll.GetAt(1))
}

type animal struct {
	name string
}

func TestLinkedList_Exists(t *testing.T) {
	ll := generic.NewLinkedList[*animal]()
	cow := &animal{name: "Cow"}
	dog := &animal{name: "Dog"}
	ll.Add(cow)
	ll.Add(dog)

	assert.True(t, ll.Exists(dog))
}

func TestLinkedList_Clear(t *testing.T) {
	ll := generic.NewLinkedList[*animal]()
	cow := &animal{name: "Cow"}
	dog := &animal{name: "Dog"}
	ll.Add(cow)
	ll.Add(dog)
	ll.Clear()

	assert.False(t, ll.Exists(dog))
	assert.Equal(t, 0, ll.GetCount())
}

func TestLinkedList_All(t *testing.T) {
	ll := generic.NewLinkedList[int16]()
	ll.Add(100)
	ll.Add(99)

	actual := slices.Collect(ll.All())

	assert.Equal(t, []int16{100, 99}, actual)
}

func TestLinkedList_GetCount(t *testing.T) {
	ll := generic.NewLinkedList[int]()
	ll.Add(100)
	ll.Add(99)

	assert.Equal(t, 2, ll.GetCount())
}
