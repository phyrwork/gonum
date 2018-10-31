package path

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/traverse"
	"gonum.org/v1/gonum/graph/internal/set"
	"gonum.org/v1/gonum/graph/internal/linear"
)

type FastestFromAt struct {
	from graph.Node
	at   int64
	nodes map[int64]struct {
		path    []graph.Node
		elapsed int64
	}
}

// From returns the starting node of the paths held by the FastestFromAt.
func (p FastestFromAt) From() graph.Node { return p.from }

type DepthFirstFastestFromAt struct {
	path linear.NodeStack
	visited set.Int64s
	now int64
}

func (d *DepthFirstFastestFromAt) FastestFromAt(g traverse.TemporalGraph, from graph.Node, at int64) FastestFromAt {
	fastest := FastestFromAt{
		from:    from,
		at:      at,
		nodes: make(map[int64]struct {
			path    []graph.Node
			elapsed int64
		}),
	}

	d.now = at
	if d.visited == nil {
		d.visited = make(set.Int64s)
	}

	var visit func (traverse.TemporalGraph, graph.Node, int64)
	visit = func (g traverse.TemporalGraph, v graph.Node, dt int64) {
		d.now += dt
		d.visited.Add(from.ID())
		vid := v.ID()
		d.path.Push(v)

		fastest.nodes[vid] = struct {
			path    []graph.Node
			elapsed int64
		}{
			append([]graph.Node{}, d.path...),
			d.now,
		}

		to := g.LinesFromAt(vid, d.now)
		for to.Next() {
			l := to.TemporalLine()
			w := l.To()
			wid := w.ID()
			ldt := l.Elapse(d.now)
			if fastest, ok := fastest.nodes[wid]; ok && d.now + ldt >= fastest.elapsed {
				continue
			}
			visit(g, w, ldt)
		}

		d.now -= dt
	}

	visit(g, from, 0)

	return fastest
}

func (d *DepthFirstFastestFromAt) Reset() {
	d.visited = nil
	d.path = nil
}