package main

import (
	"fmt"
	"strings"
)

func main() {

	mySet := NewSet(1, 2, 3)

	mySet.Add("apple")

	fmt.Println("Contains 'apple':", mySet.Contains("apple")) // true

	mySet.Remove("apple")
	fmt.Println("Contains 'apple' after removal:", mySet.Contains("apple")) // false

	fmt.Println("Set:", mySet.String())

	fmt.Println("Size of Set:", mySet.Size())

	mySet.Clear()
	fmt.Println("Set after clear:", mySet.String())
}

type Set struct {
	elements map[any]struct{}
}

func NewSet(s ...any) *Set {
	set := Set{elements: make(map[any]struct{})}
	for _, item := range s {
		set.Add(item)
	}
	return &set
}

func (s *Set) Add(element any) {
	s.elements[element] = struct{}{}
}

func (s *Set) Remove(element any) {
	delete(s.elements, element)
}

func (s *Set) Contains(element any) bool {
	_, exists := s.elements[element]
	return exists
}

func (s *Set) Size() int {
	return len(s.elements)
}

func (s *Set) Clear() {
	s.elements = make(map[any]struct{})
}

func (s *Set) String() string {
	elements := make([]string, 0, s.Size())
	for element := range s.elements {
		elements = append(elements, fmt.Sprintf("%v", element))
	}
	return "{" + strings.Join(elements, ", ") + "}"
}
