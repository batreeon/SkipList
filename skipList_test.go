package skipList

import (
	"fmt"
	"sort"
	"testing"
)

func TestNewSkipList(t *testing.T) {
	sl := NewSkipList[string, string]()
	if sl == nil {
		t.Error("skipList is nil")
	}
	if sl.len != 0 {
		t.Error("new skiplist should be empty")
	}
	if sl.level != 0 {
		t.Error("new skiplist should be empty")
	}
}

func TestInsert(t *testing.T) {
	sl := NewSkipList[string, string]()
	node := sl.Insert("foo", "bar")
	if sl.len != 1 {
		t.Error("skipList should contain 1 nodes")
	}
	if node.key != "foo" || node.value != "bar" {
		t.Errorf("expecting 'foo':'bar', got: %s:%s", node.key, node.value)
	}

	node = sl.Insert("hubei", "wuhan")
	if sl.len != 2 {
		t.Error("skipList should contain 2 nodes")
	}
	if node.key != "hubei" || node.value != "wuhan" {
		t.Errorf("expecting 'hubei':'wuhan', got: %s:%s", node.key, node.value)
	}

	if node, ok := sl.Delete("beijing"); (node != nil) || ok {
		t.Error("expecting nil and false")
	}
}

func TestSearch(t *testing.T) {
	sl := NewSkipList[string, string]()
	sl.Insert("foo", "bar")
	sl.Insert("hubei", "wuhan")
	sl.Insert("hunan", "changsha")
	sl.Insert("zhejiang", "ningbo")
	sl.Insert("zhejiang", "hangzhou")

	if node, ok := sl.Search("beijing"); (node != nil) || ok {
		t.Error("expecting nil and false")
	}
	if node, ok := sl.Search("foo"); node.value != "bar" || !ok {
		t.Errorf("expecting 'foo':'bar' and true, got: %s:%s and %t", node.key, node.value, ok)
	}
	if node, ok := sl.Search("hubei"); node.value != "wuhan" || !ok {
		t.Errorf("expecting 'hubei':'wuhan' and true, got: %s:%s and %t", node.key, node.value, ok)
	}
	if node, ok := sl.Search("hunan"); node.value != "changsha" || !ok {
		t.Errorf("expecting '':'changsha' and true, got: %s:%s and %t", node.key, node.value, ok)
	}
	if node, ok := sl.Search("zhejiang"); node.value != "hangzhou" || !ok {
		t.Errorf("expecting 'zhejiang':'hangzhou' and true, got: %s:%s and %t", node.key, node.value, ok)
	}
}

func TestDelete(t *testing.T) {
	sl := NewSkipList[string, string]()
	sl.Insert("foo", "bar")
	sl.Insert("zhejiang", "ningbo")
	sl.Insert("zhejiang", "hangzhou")

	if node, ok := sl.Delete("beijing"); (node != nil) || ok {
		t.Error("expecting nil and false")
	}
	if node, ok := sl.Delete("foo"); (node.key != "foo") || (node.value != "bar") || !ok {
		t.Errorf("expecting 'foo':'bar' and true, got: %s:%s and %t", node.key, node.value, ok)
	}
	if node, ok := sl.Delete("foo"); (node != nil) || ok {
		t.Error("expecting nil and false")
	}
	if node, ok := sl.Search("foo"); (node != nil) || ok {
		t.Error("expecting nil and false")
	}
	if node, ok := sl.Delete("zhejiang"); (node.key != "zhejiang") || (node.value != "hangzhou") || !ok {
		t.Errorf("expecting 'zhejiang':'hangzhou' and true, got: %s:%s and %t", node.key, node.value, ok)
	}
	if node, ok := sl.Search("zhejiang"); (node != nil) || ok {
		t.Error("expecting nil and false")
	}
}

func TestConcurrent(t *testing.T) {
	sl := NewSkipList[string, int]()
	ch := make(chan int)
	a := make([]int,1000)
	go func() {
		for i := 0; i < 500; i++ {
			sl.Insert(fmt.Sprint(i), i)
			node, _ := sl.Search(fmt.Sprint(i))
			ch <- node.value
		}
	}()
	go func() {
		for i := 500; i < 1000; i++ {
			sl.Insert(fmt.Sprint(i), i)
			node, _ := sl.Search(fmt.Sprint(i))
			ch <- node.value
		}
	}()

	count := 0
	for v := range ch {
		a[count] = v
		count++
		if count == 1000 {
			break
		}
	}

	if sl.len != 1000 {
		t.Error("expecting 1000 nodes")
	}

	sort.Ints(a)
	for i := 0; i < 1000; i++ {
		if a[i] != i {
			t.Error("missing value", i)
		}
	}
}