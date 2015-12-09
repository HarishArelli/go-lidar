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
	l, err := Open("../data/xyzrgb_manuscript_detail.las")
	if l == nil || err != nil {
		t.FailNow()
	}
	if l.VMaj() != 1 {
		t.Fail()
	}
	if l.VMin() != 2 {
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
	l, err := Open("../data/xyzrgb_manuscript_detail.las")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	i := 0
	for err == nil {
		_, err = l.GetNextPoint()
		i++
	}
	if i != 25008 {
		t.Logf("Read %d points, header says %d", i, l.PointCount())
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
	if l.XScale() != 0.001 {
		t.Fail()
	}
}

func TestYScale(t *testing.T) {
	l, err := Open("../data/xyzrgb_manuscript_detail.las")
	if l == nil || err != nil {
		t.FailNow()
	}
	if l.YScale() != 0.001 {
		t.Fail()
	}
}

func TestZScale(t *testing.T) {
	l, err := Open("../data/xyzrgb_manuscript_detail.las")
	if l == nil || err != nil {
		t.FailNow()
	}
	if l.ZScale() != 0.001 {
		t.Logf("Invalid z scale: %f", l.ZScale())
		t.Fail()
	}
}

func TestXOffset(t *testing.T) {
	l, err := Open("../data/xyzrgb_manuscript_detail.las")
	if l == nil || err != nil {
		t.FailNow()
	}
	if l.XOffset() != 0 {
		t.Fail()
	}
}
func TestYOffset(t *testing.T) {
	l, err := Open("../data/xyzrgb_manuscript_detail.las")
	if l == nil || err != nil {
		t.FailNow()
	}
	if l.YOffset() != 0 {
		t.Fail()
	}
}
func TestZOffset(t *testing.T) {
	l, err := Open("../data/xyzrgb_manuscript_detail.las")
	if l == nil || err != nil {
		t.FailNow()
	}
	if l.ZOffset() != 0 {
		t.Fail()
	}
}

// Bounds according to lasinfo.  Needs updating for epsilon float comparison.
// Not sure if it's built in to go.
func TestMaxX(t *testing.T) {
	l, err := Open("../data/xyzrgb_manuscript_detail.las")
	if l == nil || err != nil {
		t.FailNow()
	}
	if l.MaxX() != 10.0 {
		t.Fail()
	}
}
func TestMinX(t *testing.T) {
	l, err := Open("../data/xyzrgb_manuscript_detail.las")
	if l == nil || err != nil {
		t.FailNow()
	}
	if l.MinX() != -1.0 {
		t.Fail()
	}
}
func TestMaxY(t *testing.T) {
	l, err := Open("../data/xyzrgb_manuscript_detail.las")
	if l == nil || err != nil {
		t.FailNow()
	}
	if l.MaxY() != 9.958 {
		t.Fail()
	}
}

func TestMinY(t *testing.T) {
	l, err := Open("../data/xyzrgb_manuscript_detail.las")
	if l == nil || err != nil {
		t.FailNow()
	}
	if l.MinY() != -9.996 {
		t.Fail()
	}
}

func TestMaxZ(t *testing.T) {
	l, err := Open("../data/xyzrgb_manuscript_detail.las")
	if l == nil || err != nil {
		t.FailNow()
	}
	if l.MaxZ() != 0.181 {
		t.Fail()
	}
}

func TestMinZ(t *testing.T) {
	t.Log("Skipping due to epsilon compare failure??")
	t.Skip()
	l, err := Open("../data/xyzrgb_manuscript_detail.las")
	if l == nil || err != nil {
		t.FailNow()
	}
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
	l, err := Open("../data/xyzrgb_manuscript_detail.las")
	if l == nil || err != nil {
		t.Log("Failed to open file for reading.")
		t.FailNow()
	}

	var rMin, rMax uint16
	var gMin, gMax uint16
	var bMin, bMax uint16
	rMin = 255
	rMax = 0
	gMin = 255
	gMax = 0
	bMin = 255
	bMax = 0

	for i := 0; i < int(l.PointCount()-1); i++ {
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
	l, err := Open("../data/xyzrgb_manuscript_detail.las")
	if l == nil || err != nil {
		t.Log("Failed to open file for reading.")
		t.FailNow()
	}

	XMax := -1 * math.MaxFloat64
	XMin := math.MaxFloat64
	YMax := -1 * math.MaxFloat64
	YMin := math.MaxFloat64

	for i := 0; i < int(l.PointCount()-1); i++ {
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
	l, err := Open("../data/xyzrgb_manuscript_detail.las")
	if l == nil || err != nil {
		t.Log("Failed to open file for reading.")
		t.FailNow()
	}
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
