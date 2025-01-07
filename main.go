package main

import (
	"fmt"
	"golang-ninja/basic/cache"
)

func main() {
	inMemoryCache := cache.NewInMemoryCache()

	inMemoryCache.Set("userId", 42)
	userId := inMemoryCache.Get("userId")

	fmt.Println(userId)

	inMemoryCache.Delete("userId")
	userId = inMemoryCache.Get("userId")

	fmt.Println(userId)
}
