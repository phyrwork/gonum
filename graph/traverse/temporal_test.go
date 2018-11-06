package traverse

import (
	"gonum.org/v1/gonum/graph/temporal"
	"testing"
)

var (
	// sketchGraph is the edge stream for the arbitrary
	// graph I sketched in my workbook to demonstrate that
	// my method for earliest arrival via construction
	// worked OK.
	sketchGraph = []struct{
		F, T int64
		A, D int64
	}{
		{1, 2, 1, 0},
		{1, 3, 1, 0},
		{4, 7, 1, 0},
		{1, 4, 2, 0},
		{2, 6, 3, 0},
		{2, 5, 4, 0},
		{5, 8, 5, 0},
		{6, 8, 6, 0},
		{7, 9, 7, 0},
	}
)

func lineAt(fid, tid, a, d int64) temporal.TemporalLine {
	f := temporal.Node(fid)
	t := temporal.Node(tid)
	return temporal.TemporalLine{
		F: &f,
		T: &t,
		A: a,
		D: d,
	}
}

func TestEarliestArrivalFrom(t *testing.T) {
	tests := []struct {
		name string
		stream []struct{
			F, T int64
			A, D int64
		}
		nodes map[int64]struct{
			earliest int64
			via      []int64
		}
	}{
		{
			"sketchGraph",
			sketchGraph,
			map[int64]struct{
				earliest int64
				via      []int64
			}{
				1: {0, []int64{}},
				2: {1, []int64{1}},
				3: {1, []int64{1}},
				4: {2, []int64{1}},
				5: {4, []int64{1,2}},
				6: {3, []int64{1,2}},
				8: {5, []int64{1,2,5}},
				9: {8, []int64{1,2,6}},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})
	}
}
