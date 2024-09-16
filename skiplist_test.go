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

func TestDumpAndLoad(t *testing.T) {
	sl := NewSkipList[string, string]()
	sl.Insert("foo", "bar")
	sl.Insert("hubei", "wuhan")
	sl.Insert("hunan", "changsha")
	sl.Insert("zhejiang", "ningbo")
	sl.Insert("zhejiang", "hangzhou")

	err := sl.DumpFile("f1")
	if err != nil {
		t.Error(err)
	}
	
	sl2 := NewSkipList[string, string]()
	err = sl2.LoadFile("f1")
	if err != nil {
		t.Error(err)
	}

	for node1, node2 := sl.Front(), sl2.Front(); node1 != nil || node2 != nil; node1, node2 = node1.Next(), node2.Next() {
		if (node1 == nil && node2 != nil) || (node1 != nil && node2 == nil) {
			t.Error("node1 and node2 must be nil simultaneously")
		}
		if node1.key != node2.key || node1.value != node2.value {
			t.Errorf("expecting node1 equals to node2, but node1:{%s:%s}, node2:{%s:%s}", node1.key, node1.value, node2.key, node2.value)
		}
	}

	sl3 := NewSkipList[string, vs]()
	m := map[string]vs{
		"jiangsu": {
			"nanjing",
			"su",
			1,
		},
		"shandong": {
			"jinan",
			"lu",
			3,
		},
		"guangdong": {
			"guangzhou",
			"yue",
			4,
		},
		"zhejiang": {
			"hangzhou",
			"zhe",
			2,
		},
	}
	for k, v := range m {
		sl3.Insert(k, v)
	}

	err = sl3.DumpFile("f3")
	if err != nil {
		t.Error(err)
	}

	sl4 := NewSkipList[string, vs]()
	err = sl4.LoadFile("f3")
	if err != nil {
		t.Error(err)
	}

	for node3, node4 := sl3.Front(), sl4.Front(); node3 != nil || node4 != nil; node3, node4 = node3.Next(), node4.Next() {
		if (node3 == nil && node4 != nil) || (node3 != nil && node4 == nil) {
			t.Error("node1 and node2 must be nil simultaneously")
		}
		if node3.key != node4.key || node3.value != node4.value {
			t.Errorf("expecting node1 equals to node2, but node1:{%s:%v}, node2:{%s:%v}", node3.key, node4.value, node3.key, node4.value)
		}
	}
}

type vs struct {
	Shenghui string
	Suoxie string
	Paiming int
}