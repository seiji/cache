package cache

import (
	"fmt"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

type memClient struct {
	cli *memcache.Client
}

func newMemClient(host string, port uint32) *memClient {
	cli := memcache.New(fmt.Sprintf("%s:%d", host, port))
	cli.Timeout = 1 * time.Second
	return &memClient{cli: cli}
}

func (c *memClient) Add(item *Item) error {
	return c.cli.Add(item.memItem())
}

func (c *memClient) Decrement(key string, delta uint64) (uint64, error) {
	return c.cli.Decrement(key, delta)
}

func (c *memClient) Delete(key string) (err error) {
	err = c.cli.Delete(key)
	if err != nil && err == memcache.ErrCacheMiss {
		err = ErrCacheMiss
	}
	if err != nil {
		return
	}
	return
}

func (c *memClient) DeleteAll() error {
	return c.cli.DeleteAll()
}

func (c *memClient) FlushAll() error {
	return c.cli.FlushAll()
}

func (c *memClient) Get(key string) (item *Item, err error) {
	var i *memcache.Item
	i, err = c.cli.Get(key)
	if err != nil && err == memcache.ErrCacheMiss {
		err = ErrCacheMiss
	}
	if err != nil {
		return
	}

	item = newItemFromMemItem(i)
	return
}

func (c *memClient) GetMulti(keys []string) (items map[string]*Item, err error) {
	var m map[string]*memcache.Item
	if m, err = c.cli.GetMulti(keys); err != nil {
		return
	}
	for k, v := range m {
		items[k] = newItemFromMemItem(v)
	}

	return
}

func (c *memClient) Increment(key string, delta uint64) (uint64, error) {
	return c.cli.Increment(key, delta)
}

func (c *memClient) Replace(item *Item) error {
	return c.cli.Replace(item.memItem())
}

func (c *memClient) Set(item *Item) error {
	return c.cli.Set(item.memItem())
}

func (c *memClient) Touch(key string, seconds int32) error {
	return c.cli.Touch(key, seconds)
}

func newItemFromMemItem(item *memcache.Item) *Item {
	return &Item{
		Key:        item.Key,
		Value:      item.Value,
		Flags:      item.Flags,
		Expiration: item.Expiration,
	}
}

func (item *Item) memItem() *memcache.Item {
	return &memcache.Item{
		Key:        item.Key,
		Value:      item.Value,
		Flags:      item.Flags,
		Expiration: item.Expiration,
	}
}
