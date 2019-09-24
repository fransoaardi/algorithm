package oddsort

import (
	"reflect"
	"testing"
)

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		oddSort([]int{1, 1, 1, 2, 2, 2, 3})
	}
}

func TestOddSortRef(t *testing.T) {

	testcase := []struct {
		input  []int
		wanted []int
	}{
		{
			input:  []int{1, 1, 1, 2, 7},
			wanted: []int{1, 2, 3, 4, 7},
		},
		{
			input:  []int{1, 1, 1, 2, 2, 2, 3},
			wanted: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			input:  []int{1, 3, 3, 3, 7},
			wanted: []int{1, 3, 4, 5, 7},
		},
		{
			input:  []int{2, 2, 3, 3, 3},
			wanted: []int{2, 3, 4, 5, 6},
		},
		{
			input:  []int{3, 3, 3, 5, 13},
			wanted: []int{3, 4, 5, 7, 13},
		},
		{
			input:  []int{},
			wanted: []int{},
		},
		{
			input:  []int{1, 2, 3, 4, 5, 5, 6, 7},
			wanted: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}

	for _, val := range testcase {
		output := oddSortRef(val.input)
		if !reflect.DeepEqual(output, val.wanted) {
			t.Error(output, val.wanted)
		}
	}
}

func TestOddSort(t *testing.T) {
	testcase := []struct {
		input  []int
		wanted []int
	}{
		{
			input:  []int{1, 1, 1, 2, 7},
			wanted: []int{1, 2, 3, 4, 7},
		},
		{
			input:  []int{1, 1, 1, 2, 2, 2, 3},
			wanted: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			input:  []int{1, 3, 3, 3, 7},
			wanted: []int{1, 3, 4, 5, 7},
		},
		{
			input:  []int{2, 2, 3, 3, 3},
			wanted: []int{2, 3, 4, 5, 6},
		},
		{
			input:  []int{3, 3, 3, 5, 13},
			wanted: []int{3, 4, 5, 7, 13},
		},
		{
			input:  []int{},
			wanted: []int{},
		},
		{
			input:  []int{1, 2, 3, 4, 5, 5, 6, 7},
			wanted: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}

	for _, val := range testcase {
		output := oddSort(val.input)
		if !reflect.DeepEqual(output, val.wanted) {
			t.Error(output, val.wanted)
		}
	}
}

func TestOddSort2(t *testing.T) {
	testcase := []struct {
		input  []int
		wanted []int
	}{
		{
			input:  []int{1, 1, 1, 2, 7},
			wanted: []int{1, 2, 3, 4, 9},
		},
		{
			input:  []int{1, 1, 1, 2, 2, 2, 3},
			wanted: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			input:  []int{1, 3, 3, 3, 7},
			wanted: []int{1, 3, 4, 5, 9},
		},
		{
			input:  []int{2, 2, 3, 3, 3},
			wanted: []int{2, 3, 4, 5, 6},
		},
		{
			input:  []int{3, 3, 3, 5, 13},
			wanted: []int{3, 4, 5, 7, 15},
		},
		{
			input:  []int{},
			wanted: []int{},
		},
		{
			input:  []int{1, 2, 3, 4, 5, 5, 6, 7},
			wanted: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}

	for _, val := range testcase {
		output := oddSort2(val.input)
		if !reflect.DeepEqual(output, val.wanted) {
			t.Error(output, val.wanted)
		}
	}
}
