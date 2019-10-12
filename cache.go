// Package cachelib provides a concurrency-safe global cache
package cachelib

import (
	"encoding/json"
	"sync"
	"time"
)

type Cache interface {
	Fresh() bool
	Contents(bool) Response
	Update(interface{}) Response
}

type Response struct {
	Cached bool        `json:"cached"`
	Data   interface{} `json:"data"`
}

func (res Response) String() string {
	b, _ := json.Marshal(res)
	return string(b)
}

type Memcache struct {
	sync.Mutex
	TTL      time.Duration
	contents interface{}
	created  time.Time
}

// New returns a cache struct with no contents and a time-to-live set to the
// ttl argument passed that is a duration string
func New(ttl string) (*Memcache, error) {
	d, err := time.ParseDuration(ttl)
	if err != nil {
		return nil, err
	}
	return &Memcache{
		TTL: d,
	}, nil
}

// Fresh determines if the cache is fresh
func (c *Memcache) Fresh() bool {
	c.Lock()
	defer c.Unlock()
	if c.created.IsZero() {
		return false
	}
	if time.Since(c.created) < c.TTL {
		return true
	}
	return false
}

// Contents returns the cache and a cached bool in a Response type
func (c *Memcache) Contents(fresh bool) Response {
	c.Lock()
	defer c.Unlock()
	return Response{
		Cached: fresh,
		Data:   c.contents, // will this be a copy? No need for the lock?
	}
}

// Update updates the cache and returns it
func (c *Memcache) Update(data interface{}) Response {
	c.Lock()
	defer c.Unlock()
	c.contents = data
	c.created = time.Now()
	return Response{
		Cached: false,
		Data:   data,
	}
}
