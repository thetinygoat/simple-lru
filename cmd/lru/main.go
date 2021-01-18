package main

import (
	"fmt"

	"github.com/thetinygoat/lld/design/pkg/lru"
)

func main() {
	lru := lru.New(5)
	lru.Set("sachin", 1)
	lru.Set("saini", 2)
	lru.Set("is", 3)
	lru.Set("a", 4)
	lru.Set("dev", 5)
	fmt.Println(lru.Get("sachin"))
	lru.Set("sds", 6)
	lru.Set("ds", 7)
	lru.Set("dsds", 7)
	lru.Set("sdsdsd", 9)
}
