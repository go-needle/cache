package cache

import (
	"fmt"
	"github.com/go-needle/cache/alg"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	c := New(10, time.Duration(10)*time.Second, alg.NewLRU(1024))
	c.Add("key1", []byte("123456"))
	if v, ok := c.Get("key1"); !ok || string(v) != "123456" {
		t.Fatalf("cache hit key1=123456 failed")
	}
	time.Sleep(time.Duration(2) * time.Second)
	v, ok := c.Get("key1")
	fmt.Println(v, ok)
}

func TestCacheMore(t *testing.T) {
	c := New(10, time.Duration(10)*time.Second, alg.NewLRU(1<<32))
	for i := 0; i < 1000; i++ {
		c.Add("key"+strconv.Itoa(i), []byte(strconv.Itoa(i)))
	}
	for i := 0; i < 1000; i++ {
		num := i
		go func() {
			time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
			v, ok := c.Get("key" + strconv.Itoa(num))
			fmt.Println(v, ok)
		}()
	}
	time.Sleep(time.Duration(20) * time.Second)
}
