package las

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

type Lasf struct {
	header
	fname string
	fin   io.ReadSeeker
	index uint64
	point Pointer
}

var dbg *log.Logger

func init() {
	dbg = log.New(os.Stdout, "LASF DEBUG:", log.Ldate|log.Ltime|log.Lshortfile)
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
	fin.Seek(int64(header.PointOffset()), os.SEEK_SET)
	return &Lasf{fname: filename, fin: fin, header: header}, nil
}

var ErrInvalidFormat = errors.New("Invalid point record format")
var ErrInvalidIndex = errors.New("Invalid point record index")

func (las *Lasf) readPoint(n uint64) (Pointer, error) {
	offset := uint64(las.PointOffset()) + uint64(las.PointSize())*n
	las.fin.Seek(int64(offset), os.SEEK_SET)
	switch las.PointFormat() {
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
	las.fin.Seek(int64(las.PointOffset()), os.SEEK_SET)
	return nil
}

// GetNextPoint reads the next point in the file.  After the file is opened and
// any VLRs are read into memory, the file pointer is set to the first point.
// Each call to GetNexPoint returns the next point in the file.  This
// sequence is interupted if GetPoint is explicitly called.  This means
// GetNextPoint returns point n, then a call GetPoint(m), GetNextPoint will
// return point at m+1, not n+1.  If there is an error reading the point, or if
// we seek past the end of the points, nil and error are returned.
func (las *Lasf) GetNextPoint() (Pointer, error) {
	p, err := las.GetPoint(las.index)
	if err != nil {
		las.index = 0
	} else {
		las.index++
	}
	return p, err
}

// GetPoint fetches a specific point at index n.
func (las *Lasf) GetPoint(n uint64) (Pointer, error) {
	if n >= las.PointCount() {
		return nil, fmt.Errorf("Invalid point index %d", n)
	}
	las.index = n
	p, err := las.readPoint(n)
	if err != nil {
		return nil, err
	}
	return p, nil
}
