package consistent

import (
	"fmt"
	"sort"

	"github.com/cespare/xxhash"
	"github.com/dgryski/go-farm"
	jump "github.com/lithammer/go-jump-consistent-hash"
	"github.com/spaolacci/murmur3"
)

type hasher struct{}

func Xxhash1024(data []byte) uint64 {
	return xxhash.Sum64(data) % 4096
}

func murmurHash(data []byte) uint64 {
	return murmur3.Sum64(data) % 4096
}

func farmHash(data []byte) uint64 {
	return uint64(farm.Hash64(data)) % 4096
}

func jumpHash(data string) uint64 {
	return uint64(jump.HashString(data, 4096, jump.NewCRC64()))
}

func (hs hasher) hash_to_used(data string) uint64 { // change the hash function here
	return jumpHash(data)
}

type Hasher interface {
	hash_to_used(string) uint64
}

type Hashkey struct {
	SrcIP string
	DstIP string
}

type Consistent struct {
	hasher            Hasher             // Hash function to be used
	ring              map[uint64]string  // Store hash value -> server_id
	sortedSet         []uint64           // Store the hash value of server_id
	replicationFactor int                // Define the number of virtual node
	mapping           map[Hashkey]string // Store client_infos -> server_id
	serverList        []string           // Store the existing server
}

func NewRing(replicationFactor int) *Consistent {
	return &Consistent{
		hasher:            hasher{},
		ring:              make(map[uint64]string),
		replicationFactor: replicationFactor,
		mapping:           make(map[Hashkey]string),
	}
}

func (c *Consistent) GetRing() map[uint64]string {
	return c.ring
}

func (c *Consistent) GetHasher() Hasher {
	return c.hasher
}

func (c *Consistent) GetMapping() map[Hashkey]string {
	return c.mapping
}

func (c *Consistent) GetSortedset() []uint64 {
	return c.sortedSet
}

func (c *Consistent) AddServer(server string, num int) {
	for _, s := range c.serverList {
		if s == server {
			return
		}
	}

	// TODO:
	// Here may need to think of a better way to deal with the hash collision
	for i := 0; i < num; i++ {
		key := fmt.Sprintf("%s_%d", server, i)
		h := c.hasher.hash_to_used(key)
		if _, exists := c.ring[h]; exists {
			fmt.Printf("Avoid repetition when hash collision happen\n")
			delete(c.ring, h)
			c.delSlice(h)
			continue
		}
		c.ring[h] = server
		c.sortedSet = append(c.sortedSet, h)
	}

	sort.Slice(c.sortedSet, func(i int, j int) bool {
		return c.sortedSet[i] < c.sortedSet[j]
	})
	c.serverList = append(c.serverList, server)
	if len(c.mapping) <= 0 {
		return
	}
	c.redirectKeyFromAddServer(server)
}
func (c *Consistent) addKey(hashkey Hashkey, server string) {
	c.mapping[hashkey] = server
}

func (c *Consistent) AddKey(hashkey Hashkey) string {
	if _, exists := c.mapping[hashkey]; exists {
		return ""
	}
	server := c.MapKey(hashkey.SrcIP)
	c.addKey(hashkey, server)
	return server
}

func (c *Consistent) redirectKeyFromAddServer(server string) {
	for hashkey, _ := range c.mapping {
		if c.MapKey(hashkey.SrcIP) == server {
			c.addKey(hashkey, server)
		}
	}
}

func (c *Consistent) DelServer(server string) {
	serverfound := 0
	for i, s := range c.serverList {
		if s == server {
			c.serverList = append(c.serverList[:i], c.serverList[i+1:]...)
			serverfound = 1
			break
		}
	}
	if serverfound == 0 {
		return
	}

	// Avoid tipple hash collision in this way
	for h, s := range c.ring {
		if s == server {
			delete(c.ring, h)
			c.delSlice(h)
		}
	}

	if len(c.mapping) <= 0 && len(c.serverList) != 0 { // If in initial status, no need to redirect key
		return
	} else if len(c.mapping) <= 0 && len(c.serverList) == 0 {
		return
	}
	c.redirectKeyFromRemoveServer(server)
}

func (c *Consistent) delSlice(val uint64) { // Temporary solution for dealing with hash collision
	i := 0
	for i < len(c.sortedSet) {
		if c.sortedSet[i] == val {
			c.sortedSet = append(c.sortedSet[:i], c.sortedSet[i+1:]...)
			continue // continue instead of break because we may need to delete multiple item
		}
		i++
	}
}

func (c *Consistent) DelKey(hashkey Hashkey) string {
	if _, exists := c.mapping[hashkey]; exists {
		server := c.mapping[hashkey]
		delete(c.mapping, hashkey)
		return server
	} else {
		return ""
	}
}

func (c *Consistent) redirectKeyFromRemoveServer(server string) {
	if len(c.serverList) == 0 { // The last server is deleted, set the upServer flag to 0
		return
	}
	for k, v := range c.mapping {
		if v == server {
			// c.DelKey(k)
			c.addKey(k, c.MapKey(k.SrcIP))
		}
	}
}

func (c *Consistent) MapKey(k string) string { //O(logn)
	hash := c.hasher.hash_to_used(k)

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

func (c *Consistent) TraverseServerList() {
	for i, server := range c.serverList {
		fmt.Println("Index", i, "Server", server)
	}
}
