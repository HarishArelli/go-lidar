// Copyright (c) 2015 Boise Center Aerospace Laboratory.
// All rights reserved.  Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package las

import (
	"math"
	"os"
	"strings"
	"testing"
)

const epsilon = 0.000001

func almost(a, b float64) bool {
	return !(math.Abs(a)-math.Abs(b) > epsilon)
}

const (
	//small = "../data/xyzrgb_manuscript_detail.las"
	small = "../data/test13.las"
)

func openTest(f string, t *testing.T) *Lasf {
	l, err := Open(f)
	if l == nil || err != nil {
		t.Log(err)
		t.FailNow()
	}
	return l
}

func TestBitConv(t *testing.T) {
	i := uint8(1)
	j := convertToUInt8(i, 0, 1)
	if j != 1 {
		t.Logf("%d != 1\n", j)
		t.Fail()
	}
	i = 2
	j = convertToUInt8(i, 0, 2)
	if j != 2 {
		t.Logf("%d != 2\n", j)
		t.Fail()
	}
	j = convertToUInt8(i, 1, 0)
	if j != 0 {
		t.Logf("%d != 0\n", j)
		t.Fail()
	}
	i = 255
	var k uint8
	for k = 0; k < 8; k++ {
		j = convertToUInt8(i, k, 1)
		if j != 1 {
			t.Logf("%d != 1\n", j)
			t.Fail()
		}
	}
}

func TestOpen(t *testing.T) {
	l := openTest(small, t)
	_ = l
}

func TestOpenFail(t *testing.T) {
	l, err := Open("")
	if l != nil || err == nil {
		t.Fail()
	}
}

func TestOpenFail2(t *testing.T) {
	l, err := Open("./las_test.go")
	if l != nil || err == nil {
		t.Fail()
	}
}

func TestOpenFail3(t *testing.T) {
	l, err := Open(os.DevNull)
	if l != nil || err == nil {
		t.Fail()
	}
}

func TestSignature(t *testing.T) {
	l := openTest(small, t)
	if l.Signature() != [4]byte{'L', 'A', 'S', 'F'} {
		t.Fail()
	}
}

func TestFileSourceId(t *testing.T) {
	l := openTest(small, t)
	if l.FileSourceId() != 0 {
		t.FailNow()
	}
}

func TestGlobalEncoding(t *testing.T) {
	l := openTest(small, t)
	if l.GlobalEncoding() != 0 {
		t.FailNow()
	}
}

func TestProjectID1(t *testing.T) {
	l := openTest(small, t)
	if l.ProjectID1() != 0 {
		t.FailNow()
	}
}

func TestProjectID2(t *testing.T) {
	l := openTest(small, t)
	if l.ProjectID2() != 0 {
		t.FailNow()
	}
}

func TestProjectID3(t *testing.T) {
	l := openTest(small, t)
	if l.ProjectID3() != 0 {
		t.FailNow()
	}
}

func TestProjectID4(t *testing.T) {
	l := openTest(small, t)
	if l.ProjectID4() != [8]byte{0, 0, 0, 0, 0, 0, 0, 0} {
		t.Log(l.ProjectID4())
		t.FailNow()
	}
}

func TestVersion(t *testing.T) {
	l := openTest(small, t)
	if l.VMaj() != 1 {
		t.Fail()
	}
	if l.VMin() != 3 {
		t.Fail()
	}
}

func TestSysIdentifier(t *testing.T) {
	l := openTest(small, t)
	raw := l.SysIdentifier()
	si := string(raw[:])
	if !strings.HasPrefix(si, "LAStools (c) by Martin Isenburg") {
		t.FailNow()
	}
}

func TestGenSoftware(t *testing.T) {
	l := openTest(small, t)
	raw := l.GenSoftware()
	gs := string(raw[:])
	if !strings.HasPrefix(gs, "las2las (version 110915)") {
		t.FailNow()
	}
}

func TestCreateDOY(t *testing.T) {
	l := openTest(small, t)
	if l.CreateDOY() != 1 {
		t.Fail()
	}
}

func TestCreateYear(t *testing.T) {
	l := openTest(small, t)
	if l.CreateYear() != 1 {
		t.Fail()
	}
}

func TestHeaderSize(t *testing.T) {
	l := openTest(small, t)
	if l.HeaderSize() != 235 {
		t.Fail()
	}
}
func TestPointOffset(t *testing.T) {
	l := openTest(small, t)
	if l.PointOffset() != 909 {
		t.Fail()
	}
}
func TestVlrCount(t *testing.T) {
	l := openTest(small, t)
	if l.VlrCount() != 5 {
		t.Fail()
	}
}
func TestPointFormat(t *testing.T) {
	l := openTest(small, t)
	if l.PointFormat() != 0 {
		t.Fail()
	}
}

func TestPointCount(t *testing.T) {
	l := openTest(small, t)
	i := 0
	var err error
	for {
		_, err = l.GetNextPoint()
		if err != nil {
			break
		}
		i++
	}
	if i != 11781 {
		t.Logf("Read %d points, header says %d", i, l.PointCount())
		t.Fail()
	}
}

/*
func PointsByReturn(t *testing.T) {
	l := openTest(small, t)

}
*/

func TestXScale(t *testing.T) {
	l := openTest(small, t)
	if l.XScale() != 0.01 {
		t.Fail()
	}
}

func TestYScale(t *testing.T) {
	l := openTest(small, t)
	if l.YScale() != 0.01 {
		t.Fail()
	}
}

func TestZScale(t *testing.T) {
	l := openTest(small, t)
	if l.ZScale() != 0.01 {
		t.Logf("Invalid z scale: %f", l.ZScale())
		t.Fail()
	}
}

func TestXOffset(t *testing.T) {
	l := openTest(small, t)
	if l.XOffset() != 0 {
		t.Fail()
	}
}

func TestYOffset(t *testing.T) {
	l := openTest(small, t)
	if l.YOffset() != 0 {
		t.Fail()
	}
}

func TestZOffset(t *testing.T) {
	l := openTest(small, t)
	if l.ZOffset() != 0 {
		t.Fail()
	}
}

// Bounds according to lasinfo.  Needs updating for epsilon float comparison.
// Not sure if it's built in to go.
func TestMaxX(t *testing.T) {
	l := openTest(small, t)
	if !almost(l.MaxX(), 2484009.38) {
		t.Fail()
	}
}
func TestMinX(t *testing.T) {
	l := openTest(small, t)
	if !almost(l.MinX(), 2483569.14) {
		t.Fail()
	}
}
func TestMaxY(t *testing.T) {
	l := openTest(small, t)
	if !almost(l.MaxY(), 366616.60) {
		t.Fail()
	}
}

func TestMinY(t *testing.T) {
	l := openTest(small, t)
	if l.MinY() != 366203.87 {
		t.Fail()
	}
}

func TestMaxZ(t *testing.T) {
	l := openTest(small, t)
	if l.MaxZ() != 1581.78 {
		t.Fail()
	}
}

func TestMinZ(t *testing.T) {
	l := openTest(small, t)
	if l.MinZ() != 1480.88 {
		t.Logf("MinZ: %f", l.MinZ())
		t.Fail()
	}
}

/*
func WaveformOffset(t *testing.T) {
    l, err := Open(small)
    if l == nil || err != nil {
        t.FailNow()
    }
}
func EvlrOffset(t *testing.T) {
    l, err := Open(small)
    if l == nil || err != nil {
        t.FailNow()
    }
}
func EvlrCount(t *testing.T) {
    l, err := Open(small)
    if l == nil || err != nil {
        t.FailNow()
    }
}
*/

func TestRewind(t *testing.T) {
	l := openTest(small, t)
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

func TestOverRead(t *testing.T) {
	l := openTest(small, t)
	for {
		_, e := l.GetNextPoint()
		if e != nil {
			break
		}
	}
	p, e := l.GetNextPoint()
	if p != nil || e == nil {
		t.Fail()
	}
}

func TestFilter(t *testing.T) {
	l := openTest(small, t)
	l.SetFilter(0, 0, 0, 0)
	p, err := l.GetNextPoint()
	if p != nil || err == nil {
		t.Log("Spatial filter failed")
		t.FailNow()
	}
	l.ClearFilter()
	p, err = l.GetNextPoint()
	if p == nil || err != nil {
		t.Log("Unsetting spatial filter failed")
		t.FailNow()
	}
}

func TestFilter2(t *testing.T) {
	l := openTest(small, t)
	i := 0
	for {
		_, err := l.GetNextPoint()
		if err != nil {
			break
		}
		i++
	}
	xbuf := (l.MaxX() - l.MinX()) * 0.1
	ybuf := (l.MaxY() - l.MinY()) * 0.1
	x, y := (l.MaxX()-l.MinX())/2, (l.MaxY()-l.MinY())/2
	l.Rewind()
	l.SetFilter(x-xbuf, x+xbuf, y-ybuf, y+ybuf)
	f := 0
	for {
		_, err := l.GetNextPoint()
		if err != nil {
			break
		}
		f++
	}
	if f >= i {
		t.Fail()
	}
}

func TestQuadFilter(t *testing.T) {
	l := openTest(small, t)
	i := 0
	for {
		_, err := l.GetNextPoint()
		if err != nil {
			break
		}
		i++
	}
	xbuf := (l.MaxX() - l.MinX()) * 0.1
	ybuf := (l.MaxY() - l.MinY()) * 0.1
	x, y := (l.MaxX()-l.MinX())/2, (l.MaxY()-l.MinY())/2
	l.BuildQuadTree()
	l.SetFilter(x-xbuf, x+xbuf, y-ybuf, y+ybuf)
	f := 0
	for {
		_, err := l.GetNextPoint()
		if err != nil {
			break
		}
		f++
	}
	if f >= i {
		t.Fail()
	}
}

// Benchmarks
func BenchmarkReadAll(b *testing.B) {
	l, err := Open(small)
	if err != nil {
		return
	}
	for {
		_, err := l.GetNextPoint()
		if err != nil {
			break
		}
	}
}

const rounds = 5

func BenchmarkNormalFilter(b *testing.B) {
	l, err := Open(small)
	if err != nil {
		return
	}
	var points [rounds]int
	for i := 0; i < rounds; i++ {
		x, y := (l.MaxX()-l.MinX())/2, (l.MaxY()-l.MinY())/2
		xbuf := (l.MaxX() - l.MinX()) * 0.01 * float64(i)
		ybuf := (l.MaxY() - l.MinY()) * 0.01 * float64(i)
		l.SetFilter(x-xbuf, x+xbuf, y-ybuf, y+ybuf)
		f := 0
		for {
			_, err := l.GetNextPoint()
			if err != nil {
				break
			}
			f++
		}
		points[i] = f
	}
	for i := 1; i < rounds; i++ {
		if points[i-1] > points[i] {
			b.Fail()
		}
	}
}

func BenchmarkQuadFilter(b *testing.B) {
	l, err := Open(small)
	if err != nil {
		return
	}
	l.BuildQuadTree()
	var points [rounds]int
	for i := 0; i < rounds; i++ {
		x, y := (l.MaxX()-l.MinX())/2, (l.MaxY()-l.MinY())/2
		xbuf := (l.MaxX() - l.MinX()) * 0.01 * float64(i)
		ybuf := (l.MaxY() - l.MinY()) * 0.01 * float64(i)
		l.SetFilter(x-xbuf, x+xbuf, y-ybuf, y+ybuf)
		f := 0
		for {
			_, err := l.GetNextPoint()
			if err != nil {
				break
			}
			f++
		}
		points[i] = f
	}
	for i := 1; i < rounds; i++ {
		if points[i-1] > points[i] {
			b.Fail()
		}
	}
}
