// Package argsort implements a variant of the sort function that returns a slice of indices that would sort the array.
//
// The name comes from the popular Python numpy.Argsort function.
package argsort

import (
	"reflect"
	"sort"
)

// SortInto sorts s and populates the indices slice with the indices that would sort the input slice.
func SortInto(s sort.Interface, indices []int) {
	for i := 0; i < s.Len(); i++ {
		indices[i] = i
	}
	sort.Stable(argsorter{s: s, m: indices})
}

// Sort returns the indices that would sort s.
func Sort(s sort.Interface) []int {
	indices := make([]int, s.Len())
	SortInto(s, indices)
	return indices
}

// SortSliceInto sorts a slice and populates the indices slice with the indices that would sort the input slice.
func SortSliceInto(slice interface{}, indices []int, less func(i, j int) bool) {
	SortInto(dyn{slice, less}, indices)
}

// SortSlice return the indices that would sort a slice given a comparison function.
func SortSlice(slice interface{}, less func(i, j int) bool) []int {
	v := reflect.ValueOf(slice)
	indices := make([]int, v.Len())
	SortSliceInto(slice, indices, less)
	return indices
}

type argsorter struct {
	s sort.Interface
	m []int
}

func (a argsorter) Less(i, j int) bool { return a.s.Less(a.m[i], a.m[j]) }
func (a argsorter) Len() int           { return a.s.Len() }
func (a argsorter) Swap(i, j int)      { a.m[i], a.m[j] = a.m[j], a.m[i] }

type dyn struct {
	slice interface{}
	less  func(i, j int) bool
}

func (d dyn) Less(i, j int) bool { return d.less(i, j) }
func (d dyn) Len() int           { return reflect.ValueOf(d.slice).Len() }
func (d dyn) Swap(i, j int)      { panic("unnecessary") }
