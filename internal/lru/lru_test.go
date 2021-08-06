package lru

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var maxLRULen = 3

type cacheFeedStruct struct {
	Key, Value string
}

var cacheValues = []cacheFeedStruct{
	{Key: "testKey1", Value: "testValue1"},
	{Key: "testKey2", Value: "testValue2"},
	{Key: "testKey3", Value: "testValue3"},
	{Key: "testKey3", Value: "testValue3"},
	{Key: "testKey4", Value: "testValue4"},
	{Key: "testKey4", Value: "testValue4"},
	{Key: "testKey5", Value: "testValue5"},
	{Key: "testKey5", Value: "testValue5"},
}

func feedCache(c *LRUCache, items []cacheFeedStruct) {
	for _, item := range items {
		c.Add(item.Key, item.Value)
	}
}

func TestCacheLength(t *testing.T) {
	req := require.New(t)

	c := NewLRUCache(maxLRULen)
	feedCache(&c, cacheValues)

	req.Equal(len(c.GetInternalStore()), maxLRULen, "lru store capacity error")
}
