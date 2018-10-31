package graph

// TemporalLine is an edge in a temporal graph. A TemporalLine defines the discrete
// times between which it is available and the time cost associated with traversing
// it.
type TemporalLine interface {
	Line

	// After is the discrete time after which (inclusive) the line
	// can be traversed.
	After() int64

	// Until is the discrete time until which (inclusive) the line
	// be traversed.
	Until() int64

	// Elapse is the traversal time.
	Elapse(now int64) int64
}