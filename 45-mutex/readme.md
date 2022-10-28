# Utils

```bash
go run -race main.go # detect race conditions situations on your code while running
```

# Without caring about race conditions

```golang
package main

import (
	"fmt"
	"net/http"
)

var visitCount uint64 = 0

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		visitCount++
		w.Write([]byte(fmt.Sprintf("You're the visitor number %d!", visitCount)))
	})

	http.ListenAndServe(":3000", nil)
}
```

### first request
```bash
curl http://localhost:3000                                                                                                                                 ─╯
```

Output: You're the visitor number 1!

### after running 10k requests
```bash
ab -n 10000 -c 100 http://localhost:3000/ # make 10000 requests with 100 concurrent connections
```

### last request which we should get 10002
```bash
curl http://localhost:3000                                                                                                                                 ─╯
```

Output: You're the visitor number 9992!

---

# using mutex

```golang
package main

import (
	"fmt"
	"net/http"
	"sync"
)

var visitCount uint64 = 0

func main() {
	m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m.Lock()
		visitCount++
		m.Unlock()
		w.Write([]byte(fmt.Sprintf("You're the visitor number %d!", visitCount)))
	})

	http.ListenAndServe(":3000", nil)
}
```

### first request
```bash
curl http://localhost:3000                                                                                                                                 ─╯
```

Output: You're the visitor number 1!

### after running 10k requests
```bash
ab -n 10000 -c 100 http://localhost:3000/ # make 10000 requests with 100 concurrent connections
```

### last request which we should get 10002
```bash
curl http://localhost:3000                                                                                                                                 ─╯
```

Output: You're the visitor number 10002!

---