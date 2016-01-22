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

type header1v2 struct {
	Signature                 [4]byte // Must be "LASF"
	FileSourceId              uint16  // Reserved before 1.2
	GlobalEncoding            uint16  // Reserved before 1.2
	ProjectID1                uint32  // Optional
	ProjectID2                uint16  // Optional
	ProjectID3                uint16  // Optional
	ProjectID4                [8]byte // Optional
	VMaj                      uint8
	VMin                      uint8
	SysIdentifier             [32]byte
	GenSoftware               [32]byte
	CreateDOY                 uint16
	CreateYear                uint16
	HeaderSize                uint16
	PointOffset               uint32
	VlrCount                  uint32
	PointFormat               uint8
	PointSize                 uint16
	PointCount                uint32
	PointsByReturn            [5]uint32
	XScale, YScale, ZScale    float64
	XOffset, YOffset, ZOffset float64
	MaxX, MinX, MaxY          float64
	MinY, MaxZ, MinZ          float64
}
type header1v3 struct {
	WaveformOffset uint64
}

type header1v4 struct {
	EvlrOffset         uint64
	EvlrCount          uint32
	LongPointCount     uint64
	LongPointsByReturn [15]uint64
}

type header struct {
	header1v2
	header1v3
	header1v4
}

func readHeader(fin io.ReadSeeker) (*header, error) {
	fin.Seek(0, os.SEEK_SET)
	var h2 header1v2
	var h3 header1v3
	var h4 header1v4
	err := binary.Read(fin, binary.LittleEndian, &h2)
	if err != nil {
		return nil, err
	}
	if h2.Signature != [4]byte{'L', 'A', 'S', 'F'} {
		return nil, fmt.Errorf("invalid lasf signature: %s", string(h2.Signature[:]))
	}
	if h2.VMin >= 3 {
		err := binary.Read(fin, binary.LittleEndian, &h3)
		if err != nil {
			return nil, err
		}
	}
	if h2.VMin >= 4 {
		err := binary.Read(fin, binary.LittleEndian, &h4)
		if err != nil {
			return nil, err
		}
	}
	h := header{header1v2: h2, header1v3: h3, header1v4: h4}
	return &h, nil
}

