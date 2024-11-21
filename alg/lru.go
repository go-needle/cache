package alg

type node struct {
	val  []byte
	key  string
	prev *node
	next *node
}

// LRU is not safe for concurrent access.
type LRU struct {
	maxBytes int64
	curBytes int64
	head     *node
	tail     *node
	cache    map[string]*node
}

// NewLRU is the Constructor of LRU
func NewLRU(maxBytes int64) *LRU {
	headDummy := new(node)
	tailDummy := new(node)
	headDummy.next = tailDummy
	tailDummy.prev = headDummy
	return &LRU{
		maxBytes: maxBytes,
		head:     headDummy,
		tail:     tailDummy,
		cache:    make(map[string]*node),
	}
}

// Get look ups a key's value
func (c *LRU) Get(key string) ([]byte, bool) {
	if ele, ok := c.cache[key]; ok {
		ele.prev.next = ele.next
		c.moveToTail(ele)
		return ele.val, true
	}
	return nil, false
}

// Add adds a value to the cache.
func (c *LRU) Add(key string, value []byte) {
	if ele, ok := c.cache[key]; ok {
		ele.prev.next = ele.next
		c.moveToTail(ele)
		c.curBytes += int64(len(value) - len(ele.val))
		ele.val = value
	} else {
		ele = &node{key: key, val: value}
		c.moveToTail(ele)
		c.cache[key] = ele
		c.curBytes += int64(len(ele.key) + len(ele.val))
	}
	for c.maxBytes != 0 && c.maxBytes < c.curBytes {
		c.removeOldest()
	}
}

// Delete deletes a value from the cache.
func (c *LRU) Delete(key string) {
	if ele, ok := c.cache[key]; ok {
		ele.prev.next = ele.next
		delete(c.cache, ele.key)
		c.curBytes -= int64(len(ele.key) + len(ele.val))
	}
}

// Exist checks a value in the cache.
func (c *LRU) Exist(key string) bool {
	_, ok := c.cache[key]
	return ok
}

// removeOldest removes the oldest item
func (c *LRU) removeOldest() {
	ele := c.head.next
	if ele != nil {
		c.head.next = ele.next
		delete(c.cache, ele.key)
		c.curBytes -= int64(len(ele.key) + len(ele.val))
	}
}

// RemoveOldest removes the oldest item
func (c *LRU) moveToTail(ele *node) {
	ele.prev = c.tail.prev
	ele.next = c.tail
	c.tail.prev.next = ele
	c.tail.prev = ele
}
