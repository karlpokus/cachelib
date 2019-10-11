# cachelib
A simple cache for go apps

# usage
Create the cache
```go
import "github.com/karlpokus/cachelib"

cache := &cachelib.Memcache{TTL:time.Duration(10e9)} // 10s
```
Use the cache
```go
import "github.com/karlpokus/cachelib"

func route(cache cachelib.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if cache.Fresh() {
			fmt.Fprintf(w, "%s", cache.Contents(true))
			return
		}
	}
	// cache stale. fetch data.
	cache.Update(data)
	fmt.Fprintf(w, "%s", cache.Contents(false))
}
```

# todos
- [ ] tests
- [ ] godoc
- [ ] cache headers
- [ ] maybe a middleware opt

# license
