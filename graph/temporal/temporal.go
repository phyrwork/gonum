package temporal

import "gonum.org/v1/gonum/graph"

// Node here is a duplication of simple.Node
// to avoid needing to import both packages.

// Node is a simple graph node.
type Node int64

// ID returns the ID number of the node.
func (n Node) ID() int64 {
	return int64(n)
}

// Line is a multigraph edge.
type TemporalLine struct {
	F, T graph.Node

	UID int64

	A, D int64
}

// From returns the from-node of the line.
func (l TemporalLine) From() graph.Node { return l.F }

// To returns the to-node of the line.
func (l TemporalLine) To() graph.Node { return l.T }

// ID returns the ID of the line.
func (l TemporalLine) ID() int64 { return l.UID }

// At returns the start time of the line.
func (l TemporalLine) At() int64 { return l.A }

// During returns the traversal time of the line.
func (l TemporalLine) Duration() int64 { return l.D }