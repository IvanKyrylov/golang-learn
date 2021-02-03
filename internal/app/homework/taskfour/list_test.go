package taskfour

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("Empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("Testing List.PushFront()", func(t *testing.T) {
		l := NewList()

		lastItem := l.PushFront("iva")
		middItem := l.PushFront(20)
		firstItem := l.PushFront(false)

		require.Equal(t, firstItem, middItem.Prev)
		require.Equal(t, lastItem, middItem.Next)

		testSlice := make([]interface{}, 0, l.Len())

		for i := l.Front(); i != nil; i = i.Next {
			testSlice = append(testSlice, i.Value)
		}

		require.Equal(t, []interface{}{false, 20, "iva"}, testSlice)

	})

	t.Run("Testing List.PushBack()", func(t *testing.T) {
		l := NewList()

		firstItem := l.PushBack("iva")
		middItem := l.PushBack(20)
		lastItem := l.PushBack(false)

		require.Equal(t, lastItem, middItem.Next)
		require.Equal(t, firstItem, middItem.Prev)

		testSlice := make([]interface{}, 0, l.Len())

		for i := l.Front(); i != nil; i = i.Next {
			testSlice = append(testSlice, i.Value)
		}

		require.Equal(t, []interface{}{"iva", 20, false}, testSlice)
	})

	t.Run("Testing List.Remove()", func(t *testing.T) {

		t.Run("Nil pointer", func(t *testing.T) {
			l := NewList()
			l.Remove(nil)

			require.Equal(t, 0, l.Len())
			require.Nil(t, l.Back())
			require.Nil(t, l.Front())
		})

		t.Run("Delete in empty list", func(t *testing.T) {
			l := NewList()
			item := &ListItem{}
			l.Remove(item)

			require.Equal(t, 0, l.Len())
			require.Nil(t, l.Back())
			require.Nil(t, l.Front())
		})

		t.Run("Nil pointer ItemList.Next", func(t *testing.T) {
			l := NewList()

			l.PushBack("iva")
			l.PushBack(20)
			lastItem := l.PushBack(false)

			l.Remove(lastItem.Next)

			require.Equal(t, 3, l.Len())

			testSlice := make([]interface{}, 0, l.Len())
			for i := l.Front(); i != nil; i = i.Next {
				testSlice = append(testSlice, i.Value)
			}
			require.Equal(t, []interface{}{"iva", 20, false}, testSlice)
		})

		t.Run("Nil pointer ItemList.Prev", func(t *testing.T) {
			l := NewList()

			firstItem := l.PushBack("iva")
			l.PushBack(20)
			l.PushBack(false)
			l.Remove(firstItem.Prev)

			require.Equal(t, 3, l.Len())

			testSlice := make([]interface{}, 0, l.Len())
			for i := l.Front(); i != nil; i = i.Next {
				testSlice = append(testSlice, i.Value)
			}
			require.Equal(t, []interface{}{"iva", 20, false}, testSlice)
		})

		t.Run("Delete first item", func(t *testing.T) {
			l := NewList()

			firstItem := l.PushBack("iva")
			l.PushBack(20)
			l.PushBack(false)

			l.Remove(firstItem)

			require.Equal(t, 2, l.Len())

			testSlice := make([]interface{}, 0, l.Len())

			for i := l.Front(); i != nil; i = i.Next {
				testSlice = append(testSlice, i.Value)
			}
			require.Equal(t, []interface{}{20, false}, testSlice)
		})

		t.Run("Delete item twice", func(t *testing.T) {
			l := NewList()

			firstItem := l.PushBack("iva")
			l.PushBack(20)
			l.PushBack(false)

			l.Remove(firstItem)
			l.Remove(firstItem)

			require.Equal(t, 2, l.Len())

			testSlice := make([]interface{}, 0, l.Len())

			for i := l.Front(); i != nil; i = i.Next {
				testSlice = append(testSlice, i.Value)
			}
			require.Equal(t, []interface{}{20, false}, testSlice)

		})

		t.Run("Delete by pointer middle item", func(t *testing.T) {
			l := NewList()

			l.PushBack("iva")
			l.PushBack(20)
			lastItem := l.PushBack(false)

			l.Remove(lastItem.Prev.Prev)
			l.Remove(lastItem.Prev)
			require.Equal(t, 1, l.Len())

			testSlice := make([]interface{}, 0, l.Len())

			for i := l.Front(); i != nil; i = i.Next {
				testSlice = append(testSlice, i.Value)
			}
			require.Equal(t, []interface{}{false}, testSlice)

			require.Nil(t, lastItem.Prev)

		})

		t.Run("Delete last item in list", func(t *testing.T) {
			l := NewList()

			l.PushBack("iva")
			l.PushBack(20)
			lastItem := l.PushBack(false)

			l.Remove(lastItem.Prev.Prev)
			l.Remove(lastItem.Prev)
			l.Remove(lastItem)

			require.Equal(t, 0, l.Len())

			require.Nil(t, l.Back())
			require.Nil(t, l.Front())

		})

		t.Run("Delete middle item", func(t *testing.T) {
			l := NewList()
			firstItem := l.PushBack("iva")
			middItem := l.PushBack(20)
			lastItem := l.PushBack(false)

			l.Remove(middItem)
			require.Equal(t, 2, l.Len())

			testSlice := make([]interface{}, 0, l.Len())

			for i := l.Front(); i != nil; i = i.Next {
				testSlice = append(testSlice, i.Value)
			}
			require.Equal(t, []interface{}{"iva", false}, testSlice)

			require.Equal(t, lastItem, firstItem.Next)
			require.Equal(t, firstItem, lastItem.Prev)
		})

	})

	t.Run("Testing List.Len()", func(t *testing.T) {

		l := NewList()

		require.Equal(t, 0, l.Len())

		l.PushFront(0)
		item := l.PushBack(10)
		l.PushFront(-1)

		require.Equal(t, 3, l.Len())

		l.Remove(item)

		require.Equal(t, 2, l.Len())

	})

	t.Run("Testing List.MoveToFront()", func(t *testing.T) {

		t.Run("Nil pointer", func(t *testing.T) {
			l := NewList()

			l.MoveToFront(nil)

			require.Equal(t, 0, l.Len())
			require.Nil(t, l.Back())
			require.Nil(t, l.Front())
		})

		t.Run("Empty list", func(t *testing.T) {
			l := NewList()

			item := &ListItem{}
			l.MoveToFront(item)

			require.Equal(t, 0, l.Len())
			require.Nil(t, l.Back())
			require.Nil(t, l.Front())
		})

		t.Run("Move first item", func(t *testing.T) {
			l := NewList()

			firstItem := l.PushBack("iva")
			l.PushBack(20)
			l.PushBack(false)

			l.MoveToFront(firstItem)

			require.Equal(t, 3, l.Len())

			testSlice := make([]interface{}, 0, l.Len())
			for i := l.Front(); i != nil; i = i.Next {
				testSlice = append(testSlice, i.Value)
			}
			require.Equal(t, []interface{}{"iva", 20, false}, testSlice)

		})

		t.Run("Move last item", func(t *testing.T) {
			l := NewList()

			l.PushBack("iva")
			l.PushBack(20)
			lastItem := l.PushBack(false)

			l.MoveToFront(lastItem)

			require.Equal(t, 3, l.Len())

			testSlice := make([]interface{}, 0, l.Len())
			for i := l.Front(); i != nil; i = i.Next {
				testSlice = append(testSlice, i.Value)
			}
			require.Equal(t, []interface{}{false, "iva", 20}, testSlice)

		})

		t.Run("Move middel item", func(t *testing.T) {
			l := NewList()

			l.PushBack("iva")
			middItem := l.PushBack(20)
			l.PushBack(false)

			l.MoveToFront(middItem)

			require.Equal(t, 3, l.Len())

			testSlice := make([]interface{}, 0, l.Len())
			for i := l.Front(); i != nil; i = i.Next {
				testSlice = append(testSlice, i.Value)
			}
			require.Equal(t, []interface{}{20, "iva", false}, testSlice)

		})

	})

	t.Run("Complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
		require.Equal(t, len([]int{70, 80, 60, 40, 10, 30, 50}), l.Len())
	})
}
