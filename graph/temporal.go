package graph

type TemporalLine interface {
	Line

	// Start is the discrete time at which the line may be traversed
	At() int64

	// Duration is the discrete time which elapses while traversing the line
	Duration() int64
}