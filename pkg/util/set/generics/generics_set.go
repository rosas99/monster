package generics

import (
	"fmt"
	"sync"
)

type Set[T comparable] map[T]struct{}

func MakeSet[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) Delete(v T) {
	delete(s, v)
}

func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Iterate(f func(T)) {
	for v := range s {
		f(v)
	}
}

func main() {
	m := MakeSet[int]()
	m.Add(1)
	if m.Contains(2) {
		fmt.Println("contains 2")
	} else {
		fmt.Println("not contains 2")
	}
}

// ThreadSafeSet todo
type ThreadSafeSet[T comparable] struct {
	l sync.RWMutex
	m map[T]struct{}
}
