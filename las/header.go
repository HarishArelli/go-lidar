// Copyright (c) 2015 Boise Center Aerospace Laboratory.
// All rights reserved.  Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package las

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type header interface {
	Signature() [4]byte
	FileSourceId() uint16
	GlobalEncoding() uint16
	ProjectID1() uint32
	ProjectID2() uint16
	ProjectID3() uint16
	ProjectID4() [8]byte
	VMaj() uint8
	VMin() uint8
	SysIdentifier() [32]byte
	GenSoftware() [32]byte
	CreateDOY() uint16
	CreateYear() uint16
	HeaderSize() uint16
	PointOffset() uint32
	VlrCount() uint32
	PointFormat() uint8
	PointSize() uint16
	PointCount() uint64
	PointsByReturn() []uint64
	XScale() float64
	YScale() float64
	ZScale() float64
	XOffset() float64
	YOffset() float64
	ZOffset() float64
	MaxX() float64
	MinX() float64
	MaxY() float64
	MinY() float64
	MaxZ() float64
	MinZ() float64
	WaveformOffset() uint64
	EvlrOffset() uint64
	EvlrCount() uint32
}

const versMinorOffset = 25

func readHeader(fin io.ReadSeeker) (header, error) {
	// Check minor version
	fin.Seek(versMinorOffset, os.SEEK_SET)
	var v uint8
	err := binary.Read(fin, binary.LittleEndian, &v)
	if err != nil {
		return nil, err
	}
	fin.Seek(0, os.SEEK_SET)
	switch v {
	case 0, 1, 2:
		var h header12
		err = binary.Read(fin, binary.LittleEndian, &h)
		if err != nil {
			return nil, err
		}
		return &h, nil
	case 3:
		var h header13
		err = binary.Read(fin, binary.LittleEndian, &h)
		if err != nil {
			return nil, err
		}
		return &h, nil
	case 4:
		var h header14
		err = binary.Read(fin, binary.LittleEndian, &h)
		if err != nil {
			return nil, err
		}
		return &h, nil
	default:
		return nil, fmt.Errorf("Invalid minor version in header: %d", v)
	}
}

func newHeader(lasFmt, pFmt uint8) (header, error) {
	err := checkFmtCompat(lasFmt, pFmt)
    if err != nil {
        return nil, err
    }
    var h header
    h2 := header12{VMaj_:1, VMin_:lasFmt, PointFormat_:pFmt}
	switch lasFmt {
	case 0, 1, 2:
        h = &header12{VMaj_:1, VMin_:lasFmt, PointFormat_:pFmt}
	case 3:
        h = &header13{header12:h2}
	case 4:
        h3 := header13{header12:h2}
        h = &header14{header13:h3}
	}
    return h, nil
}

