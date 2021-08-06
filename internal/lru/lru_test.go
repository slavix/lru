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

func TestFalseOnExistKey(t *testing.T) {
	req := require.New(t)

	c := NewLRUCache(maxLRULen)

	res1 := c.Add("key1", "value1")
	req.Equal(true, res1, "not true for first putting key")

	res2 := c.Add("key1", "value1")
	req.Equal(false, res2, "not false for existing key")
}

func TestGet(t *testing.T) {
	req := require.New(t)

	c := NewLRUCache(maxLRULen)

	c.Add("key1", "value1")

	value1, ok1 := c.Get("key1")
	req.Equal("value1", value1, "c.Get() value")
	req.Equal(true, ok1, "c.Get() ok")

	value2, ok2 := c.Get("key2")
	req.Equal("", value2, "c.Get() value")
	req.Equal(false, ok2, "c.Get() ok")
}

func TestRemove(t *testing.T) {
	req := require.New(t)

	c := NewLRUCache(maxLRULen)

	c.Add("key1", "value1")

	ok1 := c.Remove("key1")
	req.Equal(true, ok1, "c.Remove() ok")

	ok2 := c.Remove("key2")
	req.Equal(false, ok2, "c.Remove() ok")
}

func TestCacheLength(t *testing.T) {
	req := require.New(t)

	c := NewLRUCache(maxLRULen)
	feedCache(&c, cacheValues)

	req.Equal(maxLRULen, len(c.GetInternalStore()), "lru store capacity error")
}

func TestOutOfLine(t *testing.T) {
	req := require.New(t)

	c := NewLRUCache(maxLRULen)
	feedCache(&c, cacheValues)

	value1, ok1 := c.Get("testKey1")
	req.Equal("", value1, "c.Get() value out of line on shouldn't existing value")
	req.Equal(false, ok1, "c.Get() ok out of line  on shouldn't existing value")

	value5, ok5 := c.Get("testKey5")
	req.Equal("testValue5", value5, "c.Get() value out of line on should existing value")
	req.Equal(true, ok5, "c.Get() ok out of line  on should existing value")
}
