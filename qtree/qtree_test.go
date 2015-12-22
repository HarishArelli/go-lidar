package qtree

import (
	"math/rand"
	"testing"
)

func eightByEight4Point(t *testing.T) *QuadTree {
	// Return an 0, 8, 0, 8 extent with points in the centroids
	q, e := New(1, 0, 8, 0, 8)
	if e != nil {
		t.Fail()
	}
	if !q.Insert(0, 2, 2) || !q.Insert(1, 6, 6) || !q.Insert(2, 2, 6) || !q.Insert(3, 6, 2) {
		t.Fail()
	}
	return q
}

func TestInvalid1(t *testing.T) {
	_, err := New(0, 0, 10, 0, 10)
	if err == nil {
		t.Fail()
	}
}

func TestInvalid2(t *testing.T) {
	_, err := New(1, 10, 0, 10, 0)
	if err == nil {
		t.Fail()
	}
}

func TestContains(t *testing.T) {
	e := envelope{0, 2, 0, 2}
	if !e.contains(point{0, 1, 1}) {
		t.Fail()
	}
	if !e.contains(point{0, 0.5, 0.5}) ||
		!e.contains(point{0, 1.5, 0.5}) ||
		!e.contains(point{0, 1.5, 1.5}) ||
		!e.contains(point{0, 0.5, 1.5}) ||
		!e.contains(point{0, 0.0, 1.0}) {
		t.FailNow()
	}
}

func TestNoContains(t *testing.T) {
	e := envelope{0, 2, 0, 2}
	if e.contains(point{0, -0.5, 0.5}) ||
		e.contains(point{0, 0.5, -0.5}) ||
		e.contains(point{0, 1.5, -0.5}) ||
		e.contains(point{0, 2.5, 0.5}) ||
		e.contains(point{0, 2.5, 1.5}) ||
		e.contains(point{0, 1.5, 2.5}) ||
		e.contains(point{0, 0.5, 2.5}) ||
		e.contains(point{0, -0.5, 1.5}) {
		t.FailNow()
	}
}

func TestCenter(t *testing.T) {
	e := envelope{0, 2, 0, 2}
	p := e.center()
	if p.x != 1.0 || p.y != 1.0 {
		t.Fail()
	}
}

func TestNorthWest(t *testing.T) {
	e := envelope{0, 2, 0, 2}
	q := e.nwQuad()
	if q.xmin != 0.0 || q.xmax != 1.0 || q.ymin != 1.0 || q.ymax != 2.0 {
		t.Fail()
	}
}

func TestNorthEast(t *testing.T) {
	e := envelope{0, 2, 0, 2}
	q := e.neQuad()
	if q.xmin != 1.0 || q.xmax != 2.0 || q.ymin != 1.0 || q.ymax != 2.0 {
		t.Fail()
	}
}

func TestSouthWest(t *testing.T) {
	e := envelope{0, 2, 0, 2}
	q := e.swQuad()
	if q.xmin != 0.0 || q.xmax != 1.0 || q.ymin != 0.0 || q.ymax != 1.0 {
		t.Fail()
	}
}

func TestSouthEast(t *testing.T) {
	e := envelope{0, 2, 0, 2}
	q := e.seQuad()
	if q.xmin != 1.0 || q.xmax != 2.0 || q.ymin != 0.0 || q.ymax != 1.0 {
		t.Fail()
	}
}

func TestNew(t *testing.T) {
	_, err := New(1, 0, 10, 0, 10)
	if err != nil {
		t.FailNow()
	}
}

func TestInsert(t *testing.T) {
	q, _ := New(1, 0, 4, 0, 4)
	if !q.Insert(0, 1, 1) {
		t.Log(q)
		t.Log("Failed to insert valid point")
		t.FailNow()
	}
	if !q.Insert(1, 2, 2) {
		t.Log("Failed to insert valid point")
		t.Log(q)
		t.Log(q.ne)
		t.Log(q.nw)
		t.Log(q.se)
		t.Log(q.sw)
		t.FailNow()
	}
	if q.Insert(2, 5, 5) {
		t.Log("Inserted invalid point")
		t.FailNow()
	}
	p := q.Query(0, 4, 0, 4)
	if len(p) != 2 {
		t.Log(p)
		t.Fail()
	}
}

func TestDepth(t *testing.T) {
	q, _ := New(1, 0, 2, 0, 2)
	if q.depth() != 1 {
		t.FailNow()
	}
	q.Insert(0, 1, 1)
	if q.depth() != 1 {
		t.FailNow()
	}
	q.Insert(1, 1, 1)
	if q.depth() != 2 {
		t.Logf("Invalid depth. Expected 2, got %d\n", q.depth())
		t.FailNow()
	}
	q.Insert(2, 1, 1)
	if q.depth() != 3 {
		t.Logf("Invalid depth. Expected 3, got %d\n", q.depth())
		t.FailNow()
	}
}

func TestQuery(t *testing.T) {
	q := eightByEight4Point(t)
	p := q.Query(0, 4, 0, 4)
	if len(p) != 1 {
		t.Logf("Didn't get expected number of records, got: %d", len(p))
		t.FailNow()
	}
	if p[0] != 0 {
		t.Fail()
	}
}

func TestQuery2(t *testing.T) {
	q := eightByEight4Point(t)
	if !q.Insert(4, 1, 2) {
		t.Fail()
	}
	p := q.Query(0, 4, 0, 4)
	if len(p) != 2 {
		t.Logf("Didn't get expected number of records, got: %d", len(p))
		t.FailNow()
	}
	if p[0] != 0 || p[1] != 4 {
		t.Fail()
	}
}

// Induce a child spawn
func TestQuery3(t *testing.T) {
	q, _ := New(1, 0, 4, 0, 4)
	i := uint64(0)
	for ; i < 5; i++ {
		if !q.Insert(i, 0.001, 0.001) {
			t.Log("Failed to insert", q)
			t.FailNow()
		}
		p := q.Query(0, 4, 0, 4)
		if len(p) != int(i+1) {
			t.Log("Invalid query:", p)
			t.FailNow()
		}
	}
}

func Benchmark1(b *testing.B) {
	side := 100.0
	q, err := New(10, 0, side, 0, side)
	if err != nil {
		b.Log(err)
		b.FailNow()
	}
	for i := 0; i < 50000; i++ {
		q.Insert(uint64(i), side*rand.Float64(), side*rand.Float64())
	}
	p := q.Query(0, 10, 0, 10)
	if len(p) < 1 {
		b.Log("Failed benchmark")
		b.Fail()
	}
}
