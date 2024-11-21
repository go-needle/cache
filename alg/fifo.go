package alg

// FIFO is not safe for concurrent access.
type FIFO struct {
	maxBytes int64
	curBytes int64
	head     *node
	tail     *node
	cache    map[string]*node
}

// NewFIFO is the Constructor of FIFO
func NewFIFO(maxBytes int64) *FIFO {
	headDummy := new(node)
	tailDummy := new(node)
	headDummy.next = tailDummy
	tailDummy.prev = headDummy
	return &FIFO{
		maxBytes: maxBytes,
		head:     headDummy,
		tail:     tailDummy,
		cache:    make(map[string]*node),
	}
}

// Get look ups a key's value
func (c *FIFO) Get(key string) ([]byte, bool) {
	if ele, ok := c.cache[key]; ok {
		return ele.val, true
	}
	return nil, false
}

// Add adds a value to the cache.
func (c *FIFO) Add(key string, value []byte) {
	if ele, ok := c.cache[key]; ok {
		c.curBytes += int64(len(value) - len(ele.val))
		ele.val = value
	} else {
		ele = &node{key: key, val: value}
		c.cache[key] = ele
		c.tail.prev.next = ele
		c.tail.prev = ele
		c.curBytes += int64(len(ele.key) + len(ele.val))
	}
	for c.maxBytes != 0 && c.maxBytes < c.curBytes {
		c.removeOldest()
	}
}

// Delete deletes a value from the cache.
func (c *FIFO) Delete(key string) {
	if ele, ok := c.cache[key]; ok {
		ele.prev.next = ele.next
		delete(c.cache, ele.key)
		c.curBytes -= int64(len(ele.key) + len(ele.val))
	}
}

// removeOldest removes the oldest item
func (c *FIFO) removeOldest() {
	ele := c.head.next
	if ele != nil {
		c.head.next = ele.next
		delete(c.cache, ele.key)
		c.curBytes -= int64(len(ele.key) + len(ele.val))
	}
}
