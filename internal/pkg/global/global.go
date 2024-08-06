package global

import "sync"

var mux sync.Mutex

func SetDB() {
	mux.Lock()
	defer mux.Unlock()
}
