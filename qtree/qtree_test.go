package qtree

import (
	"testing"
)

func eightByEight4Point() *QuadTree {
	// Return an 0, 8, 0, 8 extent with points in the centroids
	q, _ := New(1, 0, 8, 0, 8)
	q.Insert(0, 2, 2)
	q.Insert(1, 6, 6)
	q.Insert(2, 2, 6)
	q.Insert(3, 6, 2)
	return q
}
	
func TestNew(t *testing.T) {
	_, err := New(1, 0, 10, 0, 10)
	if err != nil {
		t.FailNow()
	}
}

func TestInsert(t *testing.T) {
	q, _ := New(1, 0, 10, 0, 10)
	if q.Insert(0, 0.5, 0.5) == false {
		t.FailNow()
	}
	if q.Insert(1, 9.5, 9.5) == false {
		t.FailNow()
	}
	if q.Insert(2, 20, 20) == true {
		t.FailNow()
	}
}

func TestQuery(t *testing.T) {
	q := eightByEight4Point()
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
	q := eightByEight4Point()
	q.Insert(4, 1, 2)
	p := q.Query(0, 4, 0, 4)
	if len(p) != 2 {
		t.Logf("Didn't get expected number of records, got: %d", len(p))
		t.FailNow()
	}
	if p[0] != 0 || p[1] != 4 {
		t.Fail()
	}
}

