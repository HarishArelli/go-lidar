// Copyright (c) 2015 Boise Center Aerospace Laboratory.
// All rights reserved.  Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package las

import (
	"testing"
)

const in = "../data/test13.las"

// Tests for C <-> Go compatibility

func openGo() (*Lasf, error) {
	return Open(in)
}

func openC() int {
	var fid int
	LasfOpen(in, &fid)
	return fid
}

func TestInvalidIndex(t *testing.T) {
	i := openC()
	rc := LasfReadNextPoint(i+1)
	if rc != LASF_INVALIDHANDLE {
		t.Fail()
	}
}

func TestInvalidFile(t *testing.T) {
	var i int
	rc := LasfOpen("./las.go", &i)
	if rc != LASF_CANTOPEN {
		t.Fail()
	}
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

func TestZ(t *testing.T) {
	l, _ := openGo()
	p, _ := l.GetNextPoint()
	f := openC()
	LasfReadNextPoint(f)
	var z float64
	LasfPointZ(f, &z)
	if p.Z() != z {
		t.Log(p.Z(), z)
		t.Fail()
	}
}
