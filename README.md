Go In-Memory Cache
================================

## <a href="https://github.com/roman-dubovoy/go-in-memory-cache">GitHub Repository</a>

Go In-Memory Cache helps you to cache some data in real time when app is running.

See it in action:

## Example #1

```go
package main

import (
	"fmt"
	"github.com/roman-dubovoy/go-in-memory-cache"
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
