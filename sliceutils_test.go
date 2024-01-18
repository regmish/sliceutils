package sliceutils_test

import (
	"fmt"
	"testing"

	s "github.com/shankarregmi/sliceutils"
)

func TestArrayLength(t *testing.T) {
	arr := s.Array[int]{1, 2, 3, 4, 5}
	len := arr.Length()

	if len != 5 {
		t.Errorf("Expected length to be 5, got %d", len)
	}

}

func TestArrayPush(t *testing.T) {
	arr := s.Array[int]{1, 2, 3, 4, 5}

	arr.Push(6)

	if arr.Length() != 6 {
		t.Errorf("Expected length to be 6, got %d", arr.Length())
	}

	if arr[5] != 6 {
		t.Errorf("Expected 6, got %d", arr[5])
	}
}

func TestArrayPop(t *testing.T) {
	arr := s.Array[int]{1, 2, 3, 4, 5}

	lastEl := arr.Pop()

	if lastEl != 5 {
		t.Errorf("Expected 5, got %d", lastEl)
	}

	if arr.Length() != 4 {
		t.Errorf("Expected length to be 4, got %d", arr.Length())
	}
}

func TestArrayForEach(t *testing.T) {
	arr := s.Array[int]{1, 2, 3, 4, 5}

	arr.ForEach(func(el int, idx int, arr s.Array[int]) {
		if el != arr[idx] {
			t.Errorf("Expected %d, got %d", el, arr[idx])
		}
	})
}

func TestArrayMap(t *testing.T) {
	arr := s.Array[int]{1, 2, 3, 4, 5}

	squared := arr.Map(func(el int, idx int, arr s.Array[int]) int {
		return el * el
	})

	for idx, el := range squared {
		if el != arr[idx]*arr[idx] {
			t.Errorf("Expected %d, got %d", el, arr[idx]*2)
		}
	}
}

func TestArrayFilter(t *testing.T) {
	arr := s.Array[int]{1, 2, 3, 4, 5}

	odd := arr.Filter(func(el int, idx int, arr s.Array[int]) bool {
		return el%2 != 0
	})

	if odd.Length() != 3 {
		t.Errorf("Expected length to be 3, got %d", odd.Length())
	}
}

func TestArrayIndexOf(t *testing.T) {
	arr := s.Array[int]{1, 2, 3, 4, 5}

	idx := arr.IndexOf(3)

	if idx != 2 {
		t.Errorf("Expected index to be 2, got %d", idx)
	}
}

func TestArrayIncludes(t *testing.T) {
	arr := s.Array[string]{"golang", "is", "awesome"}

	if !arr.Includes("awesome") {
		t.Errorf("Expected to include 'awesome'")
	}

	if arr.Includes("rust") {
		t.Errorf("Expected to not include 'rust'")
	}

}

func TestArraySort(t *testing.T) {
	arr := s.Array[int]{100, 3, 20, 5, 10}

	arr.Sort(func(a int, b int) bool {
		return a < b
	})

	if arr[0] != 3 {
		t.Errorf("Expected 100, got %d", arr[0])
	}

	if arr[4] != 100 {
		t.Errorf("Expected 4, got %d", arr[1])
	}
}

func TestArrayJoin(t *testing.T) {
	arr := s.Array[string]{"golang", "is", "awesome", "language"}

	str := arr.Join(" ")

	fmt.Printf("%T", str)

	if str != "golang is awesome language" {
		t.Errorf("Expected 'golang is awesome', got %s", str)
	}
}

func TestArrayReduce(t *testing.T) {
	arr := s.Array[int]{1, 2, 3, 4, 5}

	sum := arr.Reduce(func(acc int, el int, idx int, arr s.Array[int]) int {
		return acc + el
	}, 0)

	if sum != 15 {
		t.Errorf("Expected 15, got %d", sum)
	}
}
