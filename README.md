# cachelib
A simple global cache for go apps. Includes a mutex so it's concurrency safe.

[![GoDoc](https://godoc.org/github.com/karlpokus/cachelib?status.svg)](https://godoc.org/github.com/karlpokus/cachelib)

# install
```bash
$ go get github.com/karlpokus/cachelib
```

# usage
Create the cache
```go
import "github.com/karlpokus/cachelib"

cache := cachelib.New("10s")
http.ListenAndServe(":8090", route(cache))
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
		// cache stale. fetch data.
		fmt.Fprintf(w, "%s", cache.Update(data))
	}
}
```

# test
```bash
$ go test
```

# todos
- [x] tests
- [x] godoc
- [ ] cache headers
- [ ] maybe a middleware opt
- [ ] remove fresh bool

# license
MIT
