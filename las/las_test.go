package las

import (
    "math"
    "testing"
)

func TestOpen(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
    
}

func TestSignature(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
    if l.header.Signature() != [4]byte{'L', 'A', 'S', 'F'} {
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
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
    if l.header.VMaj() != 1 {
        t.Fail()
    }
    if l.header.VMin() != 2 {
        t.Fail()
    }
}
/*
func SysIdentifier(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
}
func GenSoftware(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
}
func CreateDOY(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
}
func CreateYear(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
}
func HeaderSize(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
}
func PointOffset(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
}
func VlrCount(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
}
*/
func TestPointFormat(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
    if l.header.PointFormat() != 2 {
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
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
    i := 0
    for err == nil {
        _, err = l.GetNextPoint()
        i++
    }
    if i != 25008 {
        t.Logf("Read %d points, header says %d", i, l.header.PointCount())
        t.Fail()
    }
}

/*
func PointsByReturn(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
}
*/
func TestXScale(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
    if l.header.XScale() != 0.001 {
        t.Fail()
    }
}

func TestYScale(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
    if l.header.YScale() != 0.001 {
        t.Fail()
    }
}

func TestZScale(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
    if l.header.ZScale() != 0.001 {
        t.Logf("Invalid z scale: %f", l.header.ZScale())
        t.Fail()
    }
}
/*
func XOffset(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
}
func YOffset(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
}
func ZOffset(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
}
func MaxX(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
}
func MinX(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
}
func MaxY(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
}
func MinY(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
}gg
func MaxZ(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
}
func MinZ(t *testing.T) {
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.FailNow()
    }
}
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

func TestRawExtents(t *testing.T) {
    t.SkipNow()
    l, err := Open("../data/xyzrgb_manuscript_detail.las")
    if l == nil || err != nil {
        t.Log("Failed to open file for reading.")
        t.FailNow()
    }

    XMax := -1 * math.MaxFloat64
    XMin := math.MaxFloat64
    YMax := -1 * math.MaxFloat64
    YMin := math.MaxFloat64

    for i := 0; i < int(l.PointCount() - 1); i++ {
        p, err := l.GetNextPoint()
        if err != nil {
            t.Log(err)
            t.FailNow()
        }
        x := p.X() * l.header.XScale()
        if x > XMax {
            XMax = x
        }
        if x < XMin {
            XMin = x
        }
        y := p.Y() * l.header.YScale()
        if y > YMax {
            YMax = x
        }
        if y < YMin {
            YMin = y
        }
    }
    if XMax != l.header.MaxX() {
        t.Logf("Header x max doesn't match actual(%f != %f)", XMax, l.header.MaxX())
        t.Fail()
    }
    if XMin != l.header.MinX() {
        t.Logf("Header x min doesn't match actual(%f != %f)", XMin, l.header.MinX())
        t.Fail()
    }
    if YMax != l.header.MaxY() {
        t.Logf("Header y max doesn't match actual(%f != %f)", YMax, l.header.MaxY())
        t.Fail()
    }
    if YMin != l.header.MinY() {
        t.Logf("Header y min doesn't match actual(%f != %f)", YMin, l.header.MinY())
        t.Fail()
   }
}

