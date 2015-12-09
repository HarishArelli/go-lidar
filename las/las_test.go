// Copyright (c) 2015 Boise Center Aerospace Laboratory.
// All rights reserved.  // Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.


package las

import (
	"math"
	"testing"
)

func openTest(f string, t *testing.T) *Lasf {
	l, err := Open(f)
	if l == nil || err != nil {
		t.Log(err)
		t.FailNow()
	}
	return l
}

func TestOpen(t *testing.T) {
	l := openTest("../data/xyzrgb_manuscript_detail.las", t)
	_ = l
}

func TestSignature(t *testing.T) {
	l := openTest("../data/xyzrgb_manuscript_detail.las", t)
	if l.Signature() != [4]byte{'L', 'A', 'S', 'F'} {
		t.Fail()
	}
}

/*
func TestFileSourceId(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
}
func GlobalEncoding(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
}
func ProjectID1(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
}
func ProjectID2(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
}
func ProjectID3(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
}
func ProjectID4(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
}
*/
func TestVersion(t *testing.T) {
	l := openTest("../data/xyzrgb_manuscript_detail.las", t)
	if l.VMaj() != 1 {
		t.Fail()
	}
	if l.VMin() != 2 {
		t.Fail()
	}
}

/*
func TestSysIdentifier(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
		if string(l.SysIdentifier()[:]) != "PDAL" {
			t.Logf("Invalid sys identifier : %s", l.SysIdentifier())
			t.Fail()
		}
}

func TestGenSoftware(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
		if string(l.GenSoftware()[:]) != "PDAL 9c974e46af" {
			t.Logf("Invalid generating software: %s", l.GenSoftware())
			t.Fail()
		}
}
*/
func TestCreateDOY(t *testing.T) {
	l := openTest("../data/xyzrgb_manuscript_detail.las", t)
	if l.CreateDOY() != 37 {
		t.Fail()
	}
}

func TestCreateYear(t *testing.T) {
	l := openTest("../data/xyzrgb_manuscript_detail.las", t)
	if l.CreateYear() != 2008 {
		t.Fail()
	}
}

func TestHeaderSize(t *testing.T) {
	l := openTest("../data/xyzrgb_manuscript_detail.las", t)
	if l.HeaderSize() != 227 {
		t.Fail()
	}
}
func TestPointOffset(t *testing.T) {
	l := openTest("../data/xyzrgb_manuscript_detail.las", t)
	if l.PointOffset() != 227 {
		t.Fail()
	}
}
func TestVlrCount(t *testing.T) {
	l := openTest("../data/xyzrgb_manuscript_detail.las", t)
	if l.VlrCount() != 0 {
		t.Fail()
	}
}
func TestPointFormat(t *testing.T) {
	l := openTest("../data/xyzrgb_manuscript_detail.las", t)
	if l.PointFormat() != 2 {
		t.Fail()
	}
}

/*
func PointSize(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
}
*/
func TestPointCount(t *testing.T) {
	l := openTest("../data/xyzrgb_manuscript_detail.las", t)
	i := 0
	var err error
	for err == nil {
		_, err = l.GetNextPoint()
		if err == nil {
			i++
		}
	}
	if i != 25008 {
		t.Logf("Read %d points, header says %d", i, l.PointCount())
		t.Fail()
	}
}

/*
func PointsByReturn(t *testing.T) {
	l := openTest("../data/xyzrgb_manuscript_detail.las", t)

}
*/

func TestXScale(t *testing.T) {
	l := openTest("../data/xyzrgb_manuscript_detail.las", t)
	if l.XScale() != 0.001 {
		t.Fail()
	}
}

func TestYScale(t *testing.T) {
	l := openTest("../data/xyzrgb_manuscript_detail.las", t)
	if l.YScale() != 0.001 {
		t.Fail()
	}
}

func TestZScale(t *testing.T) {
	l := openTest("../data/xyzrgb_manuscript_detail.las", t)
	if l.ZScale() != 0.001 {
		t.Logf("Invalid z scale: %f", l.ZScale())
		t.Fail()
	}
}

func TestXOffset(t *testing.T) {
	l := openTest("../data/xyzrgb_manuscript_detail.las", t)
	if l.XOffset() != 0 {
		t.Fail()
	}
}

func TestYOffset(t *testing.T) {
	l := openTest("../data/xyzrgb_manuscript_detail.las", t)
	if l.YOffset() != 0 {
		t.Fail()
	}
}

func TestZOffset(t *testing.T) {
	l := openTest("../data/xyzrgb_manuscript_detail.las", t)
	if l.ZOffset() != 0 {
		t.Fail()
	}
}

// Bounds according to lasinfo.  Needs updating for epsilon float comparison.
// Not sure if it's built in to go.
func TestMaxX(t *testing.T) {
	l := openTest("../data/xyzrgb_manuscript_detail.las", t)
	if l.MaxX() != 10.0 {
		t.Fail()
	}
}
func TestMinX(t *testing.T) {
	l := openTest("../data/xyzrgb_manuscript_detail.las", t)
	if l.MinX() != -1.0 {
		t.Fail()
	}
}
func TestMaxY(t *testing.T) {
	l := openTest("../data/xyzrgb_manuscript_detail.las", t)
	if l.MaxY() != 9.958 {
		t.Fail()
	}
}

func TestMinY(t *testing.T) {
	l := openTest("../data/xyzrgb_manuscript_detail.las", t)
	if l.MinY() != -9.996 {
		t.Fail()
	}
}

func TestMaxZ(t *testing.T) {
	l := openTest("../data/xyzrgb_manuscript_detail.las", t)
	if l.MaxZ() != 0.181 {
		t.Fail()
	}
}

func TestMinZ(t *testing.T) {
	t.Log("Skipping due to epsilon compare failure??")
	t.Skip()
	l := openTest("../data/xyzrgb_manuscript_detail.las", t)
	if l.MinZ() != -0.816 {
		t.Logf("MinZ: %f", l.MinZ())
		t.Fail()
	}
}

/*
func WaveformOffset(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
}
func EvlrOffset(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
}
func EvlrCount(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
}
*/

func TestColorRange(t *testing.T) {
	l := openTest("../data/xyzrgb_manuscript_detail.las", t)

	var rMin, rMax uint16
	var gMin, gMax uint16
	var bMin, bMax uint16
	rMin = 255
	rMax = 0
	gMin = 255
	gMax = 0
	bMin = 255
	bMax = 0

	for i := 0; i < int(l.PointCount()); i++ {
		p, err := l.GetPoint(uint64(i))
		if err != nil {
			t.Log(err)
			t.FailNow()
		}
		r, g, b := p.Red(), p.Green(), p.Blue()
		if r > rMax {
			rMax = r
		}
		if r < rMin {
			rMin = r
		}
		if g > gMax {
			gMax = g
		}
		if g < gMin {
			gMin = g
		}
		if b > bMax {
			bMax = b
		}
		if b < bMin {
			bMin = b
		}
	}

	// According to lasinfo
	if rMin != 47 || rMax != 224 || gMin != 37 || gMax != 204 || bMin != 39 || bMax != 171 {
		t.Fail()
	}
}

func TestRawExtents(t *testing.T) {
	t.Skip()
	l := openTest("../data/xyzrgb_manuscript_detail.las", t)

	XMax := -1 * math.MaxFloat64
	XMin := math.MaxFloat64
	YMax := -1 * math.MaxFloat64
	YMin := math.MaxFloat64

	for i := 0; i < int(l.PointCount()); i++ {
		p, err := l.GetNextPoint()
		if err != nil {
			t.Log(err)
			t.FailNow()
		}
		x := p.X() * l.XScale()
		if x > XMax {
			XMax = x
		}
		if x < XMin {
			XMin = x
		}
		y := p.Y() * l.YScale()
		if y > YMax {
			YMax = x
		}
		if y < YMin {
			YMin = y
		}
	}
	if XMax != l.MaxX() {
		t.Logf("Header x max doesn't match actual(%f != %f)", XMax, l.MaxX())
		t.Fail()
	}
	if XMin != l.MinX() {
		t.Logf("Header x min doesn't match actual(%f != %f)", XMin, l.MinX())
		t.Fail()
	}
	if YMax != l.MaxY() {
		t.Logf("Header y max doesn't match actual(%f != %f)", YMax, l.MaxY())
		t.Fail()
	}
	if YMin != l.MinY() {
		t.Logf("Header y min doesn't match actual(%f != %f)", YMin, l.MinY())
		t.Fail()
	}
}

func TestPoint(t *testing.T) {
	l := openTest("../data/xyzrgb_manuscript_detail.las", t)
	p, err := l.GetNextPoint()
	if err != nil {
		t.Log(err)
	}
	if p.RecordFormat() != 2 {
		t.Log("Invalid record format for point")
		t.Fail()
	}
	if p.X() < 0 {
		t.Fail()
	}
	if p.Y() < 0 {
		t.Fail()
	}
	if p.Z() < 0 {
		t.Fail()
	}
	if p.Intensity() < 0 {
		t.Fail()
	}
	if p.RetNum() < 0 {
		t.Fail()
	}
	if p.RetCount() < 0 {
		t.Fail()
	}
	if p.ScanFlag() < 0 {
		t.Fail()
	}
	if p.Edge() < 0 {
		t.Fail()
	}
	if p.Classification() < 0 {
		t.Fail()
	}
	//ClassificationString() string
	if p.ScanAngle() < 0 {
		t.Fail()
	}
	if p.UserData() < 0 {
		t.Fail()
	}
	if p.PointSourceID() < 0 {
		t.Fail()
	}
	if p.GpsTime() < 0 {
		t.Fail()
	}
	if p.Red() < 0 {
		t.Fail()
	}
	if p.Green() < 0 {
		t.Fail()
	}
	if p.Blue() < 0 {
		t.Fail()
	}
	if p.NIR() < 0 {
		t.Fail()
	}
	if p.WavePacketDesc() < 0 {
		t.Fail()
	}
	if p.WaveOffset() < 0 {
		t.Fail()
	}
	if p.WaveSize() < 0 {
		t.Fail()
	}
	if p.X_t() < 0 {
		t.Fail()
	}
	if p.Y_t() < 0 {
		t.Fail()
	}
	if p.Z_t() < 0 {
		t.Fail()
	}
}

func TestRewind(t *testing.T) {
	l := openTest("../data/xyzrgb_manuscript_detail.las", t)
	p, err := l.GetNextPoint()
	if err != nil {
		t.FailNow()
	}
	l.Rewind()
	p2, err := l.GetNextPoint()
	if p.X() != p2.X() {
		t.Fail()
	}
	if p.Y() != p2.Y() {
		t.Fail()
	}
	if p.Z() != p2.Z() {
		t.Fail()
	}
	if p.Intensity() != p2.Intensity() {
		t.Fail()
	}
	if p.RetNum() != p2.RetNum() {
		t.Fail()
	}
	if p.RetCount() != p2.RetCount() {
		t.Fail()
	}
}

func BenchmarkReadAll(b *testing.B) {
	l, err := Open("../data/xyzrgb_manuscript_detail.las")
	if err != nil {
		return
	}
	for err == nil {
		_, err := l.GetNextPoint()
		if err != nil {
			break
		}
	}
}
