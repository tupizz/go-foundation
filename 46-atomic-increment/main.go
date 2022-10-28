package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var visitCount uint64 = 0

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		visitCount++
		atomic.AddUint64(&visitCount, 1) // atomic increment / atomic operations
		w.Write([]byte(fmt.Sprintf("You're the visitor number %d!", visitCount)))
	})

	http.ListenAndServe(":3000", nil)
}
