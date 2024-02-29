package modBuffer

import "sync"

type CSBuffer struct {
	folder  string
	data    []*any
	maxSize int
	mutex   sync.Mutex
}
