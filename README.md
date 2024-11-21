<!-- markdownlint-disable MD033 MD041 -->
<div align="center">

# 🪡cache

<!-- prettier-ignore-start -->
<!-- markdownlint-disable-next-line MD036 -->
a simple cache framework for golang
<!-- prettier-ignore-end -->

<img src="https://img.shields.io/badge/golang-1.21+-blue" alt="golang">
</div>

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
	c := cache.NewLRU(1<<32, time.Duration(10)*time.Second)
	// c := cache.NewFIFO(1<<32, time.Duration(10)*time.Second)
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
```