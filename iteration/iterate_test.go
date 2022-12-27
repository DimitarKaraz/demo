package iteration

import (
	"testing"
	"reflect"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 6)
	want := "aaaaaa"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func TestSum(t *testing.T) {
	t.Run("slice", func(t *testing.T) {	
		numbers := [...]int{1,2,3,4,5,6}
		got := Sum(numbers[:]...)
		want := 21

		if got != want {
			t.Errorf("got %d, want %d, given, %v", got, want, numbers)
		}
	})
	t.Run("single number", func(t *testing.T) {	
		got := Sum(3)
		want := 3

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("single slice", func(t *testing.T) {	
		got := SumAll([]int{1,4,1})
		want := []int{6}
		
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d", got)
		}
	})
	t.Run("multiple slices", func(t *testing.T) {	
		got := SumAll([]int{1,4,1}, []int{4,3,8})
		want := []int{6,15}
		
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d", got)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	assertEquals := func(t testing.TB, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("multiple slices", func(t *testing.T) {	
		got := SumAllTails([]int{1,4,1}, []int{4,3,8})
		want := []int{5,11}
		assertEquals(t, got, want)
	})
	
	t.Run("empty slice", func(t *testing.T) {	
		got := SumAllTails([]int{}, []int{4,3,8})
		want := []int{0,11}
		assertEquals(t, got, want)
	})
}
