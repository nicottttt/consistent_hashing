package consistent

import (
	"fmt"
	"hash/fnv"
	"sort"

	"github.com/cespare/xxhash"
)

type hasher struct{}

func (hs hasher) Sum64(data []byte) uint64 {
	h := fnv.New64()
	h.Write(data)
	return h.Sum64()
}

func (hs hasher) Xxhash(data []byte) uint64 {
	return xxhash.Sum64(data)
}

func (hs hasher) Xxhash1024(data []byte) uint64 {
	return xxhash.Sum64(data) % 1024
}

type Hasher interface {
	Sum64([]byte) uint64
	Xxhash([]byte) uint64
	Xxhash1024([]byte) uint64
}

type Consistent struct {
	hasher            Hasher
	ring              map[uint64]string
	sortedSet         []uint64 // Maybe can use tree structure?
	replicationFactor int
	mapping           map[string]string // Store client_id -> server_id
}

func NewRing(replicationFactor int) *Consistent {
	return &Consistent{
		hasher:            hasher{},
		ring:              make(map[uint64]string),
		replicationFactor: replicationFactor,
		mapping:           make(map[string]string),
	}
}

func (c *Consistent) GetRing() map[uint64]string {
	return c.ring
}

func (c *Consistent) GetHasher() Hasher {
	return c.hasher
}

func (c *Consistent) AddServer(server string) {
	for i := 0; i < c.replicationFactor; i++ {
		key := []byte(fmt.Sprintf("%s%d", server, i))
		h := c.hasher.Xxhash1024(key)
		c.ring[h] = server
		c.sortedSet = append(c.sortedSet, h)
	}

	sort.Slice(c.sortedSet, func(i int, j int) bool {
		return c.sortedSet[i] < c.sortedSet[j]
	})
	if len(c.mapping) <= 0 {
		return
	}
	c.redirectKeyFromAddServer(server)
}

func (c *Consistent) addKey(key string, server string) {
	c.mapping[key] = server
}

func (c *Consistent) AddKey(key string) {
	c.addKey(key, c.MapKey(key))
}

func (c *Consistent) redirectKeyFromAddServer(server string) {
	for key := range c.mapping {
		if c.MapKey(key) == server {
			// c.DelKey(key)
			c.addKey(key, server)
		}
	}
}

func (c *Consistent) DelServer(server string) {
	for i := 0; i < c.replicationFactor; i++ { // O(n) * O(n)
		key := []byte(fmt.Sprintf("%s%d", server, i))
		h := c.hasher.Xxhash1024(key)
		delete(c.ring, h)
		c.delSlice(h)
	}

	c.redirectKeyFromRemoveServer(server) // O(mlogn)
}

func (c *Consistent) delSlice(val uint64) { // O(n)
	for i := 0; i < len(c.sortedSet); i++ {
		if c.sortedSet[i] == val {
			c.sortedSet = append(c.sortedSet[:i], c.sortedSet[i+1:]...)
			break
		}
	}
}

func (c *Consistent) DelKey(key string) {
	delete(c.mapping, key)
}

func (c *Consistent) redirectKeyFromRemoveServer(server string) {
	for k, v := range c.mapping {
		if v == server {
			// c.DelKey(k)
			c.addKey(k, c.MapKey(k))
		}
	}
}

func (c *Consistent) MapKey(k string) string { //O(logn)
	key := []byte(k)
	hash := c.hasher.Xxhash1024(key)

	idx := sort.Search(len(c.sortedSet), func(i int) bool {
		return c.sortedSet[i] >= hash
	})

	var hash_idx uint64
	// Check if the key is out of the bounds of the sortedSet
	if idx == 0 {
		hash_idx = c.sortedSet[0]
	} else if idx == len(c.sortedSet) {
		hash_idx = c.sortedSet[len(c.sortedSet)-1]
	} else {
		hash_idx = c.sortedSet[idx]
	}

	return c.ring[hash_idx]
}

func (c *Consistent) TraverseHashRing() {
	for hash, server := range c.ring {
		fmt.Println("Server", server, "hash", hash)
	}
}

func (c *Consistent) TraverseSortedSet() {
	for i, hash := range c.sortedSet {
		fmt.Println("Index", i, "hash", hash)
	}
}

func (c *Consistent) TraverseMapping() {
	for key, server := range c.mapping {
		fmt.Println("Key", key, "Server", server)
	}
}
