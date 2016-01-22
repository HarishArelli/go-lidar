// Copyright (c) 2015 Boise Center Aerospace Laboratory.
// All rights reserved.  Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package las implements read access to LASF style lidar files.  All header
// versions (0, 1, 2, 3, 4) are supported.
package las

import (
	"errors"
	"fmt"
	"io"
	"math"
	"os"

	"github.com/bcal-lidar/go-lidar/qtree"
)

type filter struct {
	xmin, xmax, ymin, ymax float64
}

func (f filter) contains(x, y float64) bool {
	c := x > f.xmin && x < f.xmax && y > f.ymin && y < f.ymax
	return c
}

type Lasf struct {
	*header
	fname string
	fin   io.ReadSeeker
	index uint64
	point Pointer
	filter
	qt   *qtree.QuadTree
	qids []uint64
	qid  uint64
}

// Check and make sure this is correct...
func convertToUInt8(uval, start, length uint8) uint8 {
	c := uint8(0)
	for i, j := start, 0; j < int(length); i, j = i+1, j+1 {
		if uval&(1<<i)>>i > 0 {
			c += uint8(math.Pow(2, float64(j)))
		}
	}
	return c
}

// Open attempts to open filename and read the LASF header.  If the file is not
// a valid LASF file, or it cannot be opened, nil and the associated error is
// returned.
func Open(filename string) (*Lasf, error) {
	fin, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	header, err := readHeader(fin)
	if err != nil {
		return nil, err
	}
	// Seek to the start of the points
	fin.Seek(int64(header.PointOffset), os.SEEK_SET)
	filt := filter{-1 * math.MaxFloat64, math.MaxFloat64, -1 * math.MaxFloat64, math.MaxFloat64}
	l := Lasf{fname: filename, fin: fin, header: header, filter: filt}
	l.readVlrs()
	return &l, nil
}

var ErrInvalidFormat = errors.New("Invalid point record format")
var ErrInvalidIndex = errors.New("Invalid point record index")

func (las *Lasf) readPoint(n uint64) (Pointer, error) {
	offset := uint64(las.PointOffset) + uint64(las.PointSize)*n
	las.fin.Seek(int64(offset), os.SEEK_SET)
	switch las.PointFormat {
	case 0:
		return las.readPointFormat0()
	case 1:
		return las.readPointFormat1()
	case 2:
		return las.readPointFormat2()
	case 3:
		return las.readPointFormat3()
	case 4:
		return las.readPointFormat4()
	case 5:
		return las.readPointFormat5()
	case 6:
		return las.readPointFormat6()
	case 7:
		return las.readPointFormat7()
	case 8:
		return las.readPointFormat8()
	case 9:
		return las.readPointFormat9()
	case 10:
		return las.readPointFormat10()
	default:
		return nil, ErrInvalidFormat
	}
}

// Rewind resets the the point index to the first point in the file
func (las *Lasf) Rewind() error {
	las.index = 0
	las.fin.Seek(int64(las.PointOffset), os.SEEK_SET)
	return nil
}

// GetNextPoint reads the next point in the file.  After the file is opened and
// any VLRs are read into memory, the file pointer is set to the first point.
// Each call to GetNextPoint returns the next point in the file.  This
// sequence is interupted if GetPoint is explicitly called.  This means
// GetNextPoint returns point n, then a call GetPoint(m), GetNextPoint will
// return point at m+1, not n+1.  If there is an error reading the point, or if
// we seek past the end of the points, nil and error are returned.
func (las *Lasf) GetNextPoint() (Pointer, error) {
	i := uint64(0)
	for {
		if las.qt != nil && las.qids != nil {
			if int(las.qid) >= len(las.qids) {
				return nil, ErrInvalidIndex
			}
			i = las.qids[las.qid]
			if i >= uint64(las.PointCount()) {
				las.qid = 0
				return nil, ErrInvalidIndex
			}
			las.qid++
		} else {
			i = las.index
			las.index++
		}
		p, err := las.GetPoint(i)
		if err != nil {
			return p, err
		}
		if las.contains(p.X()*las.XScale, p.Y()*las.YScale) {
			return p, nil
		}
	}
}

// GetPoint fetches a specific point at index n.
func (las *Lasf) GetPoint(n uint64) (Pointer, error) {
	if n >= uint64(las.PointCount()) {
		return nil, fmt.Errorf("Invalid point index %d", n)
	}
	p, err := las.readPoint(n)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// SetFilter applies a minimum bounding rectangle spatial filter on points
// returned by GetNextPoint.  GetPoint is unaffected by the filter.  If a
// QuadTree has been built, it is used to potentially speed up access.
func (las *Lasf) SetFilter(xmin, xmax, ymin, ymax float64) {
	las.xmin = xmin
	las.xmax = xmax
	las.ymin = ymin
	las.ymax = ymax
	las.Rewind()
	if las.qt != nil {
		las.qids = las.qt.Query(xmin, xmax, ymin, ymax)
	}
	las.qid = 0
}

// ClearFilter clears the spatial filter.  Rewind() is called and sequential
// reading is reset.
func (las *Lasf) ClearFilter() {
	las.SetFilter(-1*math.MaxFloat64, math.MaxFloat64, -1*math.MaxFloat64, math.MaxFloat64)
	las.Rewind()
	las.qids = nil
	las.qid = 0
}

// BuildQuadTree creates a simple spatial index to potentially speed up
// filtered reads using GetNextPoint.
func (las *Lasf) BuildQuadTree() {
	las.ClearFilter()
	n := uint64(float64(las.PointCount() / 10.0))
	qt, err := qtree.New(n, las.MinX, las.MaxX, las.MinY, las.MaxY)
	if err != nil {
		return
	}
	i := uint64(0)
	for {
		p, err := las.GetNextPoint()
		if err != nil {
			break
		}
		qt.Insert(i, p.X()*las.XScale, p.Y()*las.YScale)
		i++
	}
	las.qid = 0
	las.qt = qt
}
