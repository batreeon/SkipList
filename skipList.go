package skipList

import (
	"math/rand"
	"sync"
	. "golang.org/x/exp/constraints"
)

const (
	p = 0.5
	maxLevel = 32
)

type skipList[K Ordered, V any] struct {
	mutex sync.RWMutex
	// 跳表的头节点（哑节点）
	header *skipNode[K,V]
	// 跳表的节点个数
	len int
	// 跳表的最大level数
	level int
} 

type skipNode[K Ordered, V any] struct {
	// 跳表节点根据key值排序
	key K
	value V
	// 索引
	next []*skipNode[K, V]
}

func NewSkipList[K Ordered, V any] () *skipList[K, V] {
	var k K
	var v V
	return &skipList[K, V]{
		header: NewSkipNode(k, v, maxLevel),
		len: 0,
		level: 0,
	}
}

func NewSkipNode[K Ordered, V any] (k K, v V, level int) *skipNode[K, V] {
	return &skipNode[K, V]{
		key: k,
		value: v,
		next: make([]*skipNode[K, V], level),
	}
}

func (node skipNode[K, V]) Key() K {
	return node.key
}

func (node skipNode[K, V]) Value() V {
	return node.value
}

func (node skipNode[K, V]) Next() *skipNode[K, V] {
	return node.next[0]
}

func (sl *skipList[K, V]) Front() *skipNode[K, V] {
	return sl.header.next[0]
}

func randomLevel() int {
	level := 1
	for rand.Float64() < 0.5 && level < maxLevel {
		level++
	}
	return level
}

func (sl *skipList[K, V]) Search(k K) (*skipNode[K, V], bool) {
	sl.mutex.RLock()
	defer sl.mutex.RUnlock()
	x := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for x.next[i] != nil && x.next[i].key < k {
			x = x.next[i]
		}
	}
	x = x.next[0]
	if x != nil && x.key == k {
		return x, true
	}
	return nil, false
}

func (sl *skipList[K, V]) Insert(k K, v V) *skipNode[K, V] {
	sl.mutex.Lock()
	defer sl.mutex.Unlock()
	update := make([]*skipNode[K, V], maxLevel)
	x := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for x.next[i] != nil && x.next[i].key < k {
			x = x.next[i]
		}
		update[i] = x
	}
	x = x.next[0]

	// key已经存在
	if x != nil && x.key == k {
		x.value = v
		return x
	}
	// key不存在，需要新建节点
	newLevel := randomLevel()
	if newLevel > sl.level {
		for i := sl.level; i < newLevel; i++ {
			update[i] = sl.header
		}
		sl.level = newLevel
	}
	newNode := NewSkipNode(k, v, newLevel)
	for i := 0 ; i < newLevel; i++ {
		newNode.next[i] = update[i].next[i]
		update[i].next[i] = newNode
	}

	sl.len++
	return newNode
}

func (sl *skipList[K, V]) Delete(k K) (*skipNode[K, V], bool) {
	sl.mutex.Lock()
	defer sl.mutex.Unlock()
	update := make([]*skipNode[K, V], maxLevel)
	x := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for x.next[i] != nil && x.next[i].key < k {
			x = x.next[i]
		}
		update[i] = x
	}
	x = x.next[0]

	// 判断是否存在想要删除的节点 
	if x == nil || x.key != k {
		return nil, false
	}

	for i := 0; i < sl.level; i++ {
		if update[i].next[i].key != k {
			break
		}
		update[i].next[i] = x.next[i]
	}

	return x, true
}

