package las

import (
	"testing"
)

const in = "../data/xyzrgb_manuscript_detail.las"

// Tests for C <-> Go compatibility

func openGo() (*Lasf, error) {
	return Open(in)
}

func openC() int {
	var fid int
	LasfOpen(in, &fid)
	return fid
}

func TestX(t *testing.T) {
	l, _ := openGo()
	p, _ := l.GetNextPoint()
	f := openC()
	LasfReadNextPoint(f)
	var x float64
	LasfPointX(f, &x)
	if p.X() != x {
		t.Log(p.X(), x)
		t.Fail()
	}
}

func TestY(t *testing.T) {
	l, _ := openGo()
	p, _ := l.GetNextPoint()
	f := openC()
	LasfReadNextPoint(f)
	var y float64
	LasfPointY(f, &y)
	if p.Y() != y {
		t.Log(p.Y(), y)
		t.Fail()
	}
}