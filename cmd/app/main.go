package main

import (
	"github/slavix/lru/internal/lru"
	"log"
)

func main() {
	log.Println("test")

	c := lru.NewLRUCache(2)

	c.Add("key1", "value1")
	c.Add("key2", "value2")
	c.Add("key3", "value3")

	c.Get("key1")

	c.Remove("key3")
}
