// Copyright ©2014 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package traverse

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/internal/set"
)

type TemporalGraph interface {
	LinesFromAt(id, t int64) graph.TemporalLines
}

type TemporalDepthFirst struct {
	LineFilter func(graph.TemporalLine) bool
	Visit      func(u, v graph.Node)
	visited    set.Int64s
	now        int64
}

func (d *TemporalDepthFirst) Walk(g TemporalGraph, from graph.Node, at int64, until func(graph.Node) bool) graph.Node {
	d.now = at
	if d.visited == nil {
		d.visited = make(set.Int64s)
	}

	var visit func (TemporalGraph, graph.Node, int64) graph.Node
	visit = func (g TemporalGraph, v graph.Node, dt int64) graph.Node {
		d.now += dt
		d.visited.Add(from.ID())
		if until != nil && until(v) {
			return v
		}
		vid := v.ID()
		to := g.LinesFromAt(vid, d.now)
		for to.Next() {
			l := to.TemporalLine()
			w := l.To()
			wid := w.ID()
			if d.LineFilter != nil && !d.LineFilter(l) {
				continue
			}
			if d.visited.Has(wid) {
				continue
			}
			if d.Visit != nil {
				d.Visit(v, w)
			}
			visit(g, w, l.Elapse(d.now))
		}
		d.now -= dt
		return nil
	}

	return visit(g, from, 0)
}

func (d *TemporalDepthFirst) Reset() {
	d.visited = nil
}