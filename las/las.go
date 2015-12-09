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
	fname  string
	fin    io.ReadSeeker
	header Header
	index  uint64
	point  Pointer
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

func (las *Lasf) open() error {
	fin, err := os.Open(las.fname)
	if err != nil {
		return err
	}
	las.fin = fin
	header, err := las.readHeader()
	if err != nil {
		return err
	}
	las.header = header
	return nil
}

func Open(filename string) (*Lasf, error) {
	var las Lasf
	las.fname = filename
	err := las.open()
	if err != nil {
		return nil, err
	}
	return &las, nil
}

var ErrInvalidFormat = errors.New("Invalid point record format")
var ErrInvalidIndex = errors.New("Invalid point record index")

func (las *Lasf) readPoint(n uint64) (Pointer, error) {
	offset := uint64(las.header.PointOffset()) + uint64(las.header.PointSize())*n
	las.fin.Seek(int64(offset), os.SEEK_SET)
	switch las.header.PointFormat() {
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

func (las *Lasf) Rewind() error {
	las.index = 0
	las.fin.Seek(0, os.SEEK_SET)
	return nil
}

func (las *Lasf) GetNextPoint() (Pointer, error) {
	las.index++
	return las.GetPoint(las.index)
}

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

func (las *Lasf) PointCount() uint64 {
	return las.header.PointCount()
}
