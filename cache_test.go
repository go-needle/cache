package cache

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	c := NewFIFO(1024, time.Duration(4)*time.Second)
	//c := NewLRU(1024, time.Duration(4)*time.Second)
	c.Add("key1", []byte("123456"))
	if v, ok := c.Get("key1"); !ok || v.String() != "123456" {
		t.Fatalf("cache hit key1=123456 failed")
	}
	time.Sleep(time.Duration(2) * time.Second)
	v, ok := c.Get("key1")
	fmt.Println(v, ok)
}

func TestCacheMoreLRU(t *testing.T) {
	c := NewLRU(1<<32, time.Duration(10)*time.Second)
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

func TestCacheMoreFIFO(t *testing.T) {
	c := NewFIFO(1<<32, time.Duration(10)*time.Second)
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
