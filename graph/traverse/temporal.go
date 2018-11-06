package traverse

import (
	"gonum.org/v1/gonum/graph"
	"log"
	"gonum.org/v1/gonum/graph/temporal"
)

type TemporalStream interface {
	LineStreamAt(t int64) graph.TemporalLines
}

type Earliest struct {
	from  graph.Node
	at    int64
	until int64
	nodes map[int64]struct{
		earliest int64
		via      []int64
	}
}

func (e *Earliest) set(v graph.Node, t int64, p []int64) {
	e.nodes[v.ID()] = struct{
		earliest int64
		via      []int64
	}{
		t,
		p,
	}
}

func (e *Earliest) From() graph.Node { return e.from }

func (e *Earliest) At() int64 { return e.at }

func (e *Earliest) Until() int64 { return e.until }

func (e *Earliest) To(u interface{}) (path []graph.Node, duration int64) {
	var uid int64
	switch t := u.(type) {
	case graph.Node:
		uid = t.ID()
	case int64:
		uid = t
	default:
		log.Panicln("not supported")
	}
	if eu, ok := e.nodes[uid]; !ok {
		return nil, -1
	} else {
		duration = eu.earliest
		for _, vid := range eu.via {
			v := temporal.Node(vid)
			path = append(path, &v)
		}
		u := temporal.Node(uid)
		path = append(path, &u)
		return
	}
}

func EarliestArrivalFrom(s graph.TemporalLines, from graph.Node, at int64, until int64) Earliest {
	earliest := Earliest{
		from:  from,
		at:    at,
		until: until,
		nodes: make(map[int64]struct{
			earliest int64
			via      []int64
		}),
	}
	earliest.set(from, at, []int64{})
	for s.Next() {
		l := s.TemporalLine()
		u := l.From()
		uid := u.ID()
		eu, ok := earliest.nodes[uid]
		tl := l.At()
		dtl := l.Duration()
		if !ok {
			// Ignore edge from unseen node
			continue
		}
		if tl + dtl <= until && tl >= eu.earliest {
			v := l.To()
			vid := v.ID()
			ev, ok := earliest.nodes[vid]
			if !ok || tl + dtl < ev.earliest {
				earliest.set(v, tl + dtl, append(eu.via, uid))
			}
		} else if tl >= until {
			break
		}
	}
	return earliest
}