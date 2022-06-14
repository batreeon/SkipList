package skipList

import (
	"fmt"
	"testing"
)

func BenchmarkInsert1(b *testing.B) {
	sl := NewSkipList[string, int]()

	for i := 0; i < b.N; i++ {
		sl.Insert(fmt.Sprint(i), i)
	}
}

func BenchmarkInsert2(b *testing.B) {
	sl2 := NewSkipList[string, vs]()

	for i := 0; i < b.N; i++ {
		sl2.Insert(fmt.Sprint(i), vs{fmt.Sprint(i),fmt.Sprint(i),i,})
	}
}


func BenchmarkSearch(b *testing.B) {
	sl2 := NewSkipList[string, vs]()

	for i := 0; i < 100000; i++ {
		sl2.Insert(fmt.Sprint(i), vs{fmt.Sprint(i),fmt.Sprint(i),i,})
	}

	for i := 0; i < b.N; i++ {
		sl2.Search(fmt.Sprint(i%100000))
	}
}