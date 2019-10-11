package cachelib

import (
	"encoding/json"
	"sync"
	"time"
)

type Cache interface {
	Fresh() bool
	Contents(bool) Response
	Update(interface{})
}

type Response struct {
	Cached bool
	Data   interface{}
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

// Contents returns the cached contents and a Cached field set to the
// passed fresh argument
func (c *Memcache) Contents(fresh bool) Response {
	c.Lock()
	defer c.Unlock()
	return Response{
		Cached: fresh,
		Data:   c.contents,
	}
}

// Update updates the cache contents
func (c *Memcache) Update(data interface{}) {
	c.Lock()
	defer c.Unlock()
	c.contents = data
	c.created = time.Now()
}
