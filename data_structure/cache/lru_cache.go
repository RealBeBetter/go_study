package cache

import (
	"container/list"
)

const defaultCapacity = 10

type LRUCache struct {
	KeyToNodeMap map[int]*list.Element
	LinkedList   *list.List
	Size         int
	Capacity     int
}

type Node struct {
	Key   int
	Value int
}

func Constructor(capacity int) LRUCache {
	lruCache := new(LRUCache)
	lruCache.Size = 0
	lruCache.Capacity = capacity
	if capacity <= 0 {
		lruCache.Capacity = defaultCapacity
	}
	lruCache.KeyToNodeMap = make(map[int]*list.Element)
	lruCache.LinkedList = list.New()

	return *lruCache
}

func (cache *LRUCache) Get(key int) int {
	if valueNode, ok := cache.KeyToNodeMap[key]; ok {
		// 移动到队尾
		cache.LinkedList.MoveToBack(valueNode)
		return valueNode.Value.(Node).Value
	}

	// 不存在则返回 -1
	return -1
}

func (cache *LRUCache) Put(key int, value int) {
	// 添加的时候存在值
	if currentNode, ok := cache.KeyToNodeMap[key]; ok {
		cache.LinkedList.Remove(currentNode)
		// 将添加的节点放到队尾
		cache.KeyToNodeMap[key] = cache.LinkedList.PushBack(Node{Key: key, Value: value})
		return
	}

	// 不存在需要判断 size 是否符合要求
	if cache.Capacity == cache.Size {
		// 删除最久未使用的队头，map 中存储的 Key，所以需要自定义 Node
		delete(cache.KeyToNodeMap, cache.LinkedList.Front().Value.(Node).Key)
		cache.LinkedList.Remove(cache.LinkedList.Front())
		cache.Size--
	}

	cache.LinkedList.PushBack(Node{Key: key, Value: value})
	cache.KeyToNodeMap[key] = cache.LinkedList.Back()
	cache.Size++
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
