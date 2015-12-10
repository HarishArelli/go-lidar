// Copyright (c) 2015 Boise Center Aerospace Laboratory.
// All rights reserved.  Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package qtree implements a very simple QuadTree for spatial searching.
package qtree

import (
	"fmt"
)

type envelope struct {
	xmin, xmax, ymin, ymax float64
}

type point struct {
	i    uint64
	x, y float64
}

func (e *envelope) contains(p point) bool {
	c := p.x < e.xmax && p.x > e.xmin && p.y < e.ymax && p.y > e.ymin
	return c
}

func (e *envelope) intersects(other *envelope) bool {
	c := e.contains(point{x: other.xmin, y: other.ymax}) ||
		e.contains(point{x: other.xmax, y: other.ymax}) ||
		e.contains(point{x: other.xmax, y: other.ymin}) ||
		e.contains(point{x: other.xmin, y: other.ymin})
	return c
}

func (e *envelope) center() point {
	x := (e.xmax - e.xmin) / 2
	y := (e.ymax - e.ymin) / 2
	return point{x: x, y: y}
}

func (e *envelope) nwQuad() envelope {
	c := e.center()
	return envelope{xmin: e.xmin, xmax: c.x, ymin: c.y, ymax: e.ymax}
}

func (e *envelope) neQuad() envelope {
	c := e.center()
	return envelope{xmin: c.x, xmax: e.xmax, ymin: c.y, ymax: e.ymax}
}

func (e *envelope) swQuad() envelope {
	c := e.center()
	return envelope{xmin: e.xmin, xmax: c.x, ymin: e.ymin, ymax: c.y}
}

func (e *envelope) seQuad() envelope {
	c := e.center()
	return envelope{xmin: c.x, xmax: e.xmax, ymin: e.ymin, ymax: c.y}
}

type QuadTree struct {
	capacity uint64
	length   uint64
	points   []point
	envelope
	nw, ne, sw, se *QuadTree
}

// New returns a usable QuadTree with a capacity at each node of capacity.  It
// covers the extent defined by xmin, xmax, ymin, ymax
func New(capacity uint64, xmin, xmax, ymin, ymax float64) (*QuadTree, error) {
	if capacity < 1 {
		return nil, fmt.Errorf("Capacity must be greater than 0")
	}
	env := envelope{xmin, xmax, ymin, ymax}
	return &QuadTree{capacity: capacity, envelope: env}, nil
}

// Insert adds a point to the QuadTree.  The point consists of an x, y
// coordinate and an id, which is assumed to be unique.
func (q *QuadTree) Insert(i uint64, x, y float64) bool {
	p := point{i, x, y}
	if !q.contains(p) {
		return false
	}
	if q.length < q.capacity {
		q.points = append(q.points, p)
		return true
	}

	if q.nw == nil {
		env := q.nwQuad()
		q.nw = &QuadTree{capacity: q.capacity, envelope: env}
		env = q.neQuad()
		q.ne = &QuadTree{capacity: q.capacity, envelope: env}
		env = q.swQuad()
		q.sw = &QuadTree{capacity: q.capacity, envelope: env}
		env = q.seQuad()
		q.se = &QuadTree{capacity: q.capacity, envelope: env}
	}
	if q.nw.Insert(i, p.x, p.y) == true {
		return true
	}
	if q.ne.Insert(i, p.x, p.y) == true {
		return true
	}
	if q.sw.Insert(i, p.x, p.y) == true {
		return true
	}
	if q.se.Insert(i, p.x, p.y) == true {
		return true
	}
	panic("NEVER") // Point not in envelope?
}

// Query returns a list of unique ids of points that lie within pages that
// intersect the minimum bounding rectangle defined by xmin, xmax, ymin, ymax.
// Note that some points will not intersect the MBR exactly.
func (q *QuadTree) Query(xmin, xmax, ymin, ymax float64) []uint64 {
	var points []uint64
	qenv := envelope{xmin, xmax, ymin, ymax}
	if !q.intersects(&qenv) {
		return points
	}
	for _, p := range q.points {
		if qenv.contains(p) {
			points = append(points, p.i)
		}
	}
	if q.nw == nil {
		return points
	}
	points = append(points, q.nw.Query(xmin, xmax, ymin, ymax)...)
	points = append(points, q.ne.Query(xmin, xmax, ymin, ymax)...)
	points = append(points, q.sw.Query(xmin, xmax, ymin, ymax)...)
	points = append(points, q.se.Query(xmin, xmax, ymin, ymax)...)
	return points
}
