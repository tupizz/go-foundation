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
