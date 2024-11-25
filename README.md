<!-- markdownlint-disable MD033 MD041 -->
<div align="center">

# 🪡cache

<!-- prettier-ignore-start -->
<!-- markdownlint-disable-next-line MD036 -->
a simple cache framework for golang
<!-- prettier-ignore-end -->

<img src="https://img.shields.io/badge/golang-1.21+-blue" alt="golang">
</div>

## introduction
This is a cache framework implemented by Golang. This framework implements the FIFO and LRU algorithms for managing cache and is thread safe. The algorithm needs to be selected based on the actual situation, and using FIFO in an environment with more reads and less writes will result in higher performance.

## installing
Select the version to install

`go get github.com/go-needle/cache@version`

If you have already get , you may need to update to the latest version

`go get -u github.com/go-needle/cache`


## quickly start
```golang
package main

import (
	"fmt"
	"github.com/go-needle/cache"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	c := cache.NewLRU(1024, time.Duration(10)*time.Second)
	//c := cache.NewFIFO(1<<32, time.Duration(10)*time.Second)
	for i := 0; i < 1000; i++ {
		num := i
		go func() {
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
			c.Add("key"+strconv.Itoa(num), []byte(strconv.Itoa(num)))
		}()
	}
	for i := 0; i < 1000; i++ {
		num := i
		go func() {
			time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
			v, ok := c.Get("key" + strconv.Itoa(num))
			fmt.Println(v.String(), ok)
		}()
	}
	time.Sleep(time.Duration(20) * time.Second)
}
```
