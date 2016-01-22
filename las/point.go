// Copyright (c) 2015 Boise Center Aerospace Laboratory.
// All rights reserved.  Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package las

import (
	"encoding/binary"
	"os"
)

// Represents a fully capable LASF point.  The interface handles any field
// available in *any* point format from the LASF 1.4 spec.  If a field is not
// supported, 0 or nil is returned.  This may possible change to return an
// error as well.

type Pointer interface {
	RecordFormat() uint8
	X() float64
	Y() float64
	Z() float64
	Intensity() uint16
	RetNum() uint8
	RetCount() uint8
	ScanFlag() uint8
	Edge() uint8
	Classification() uint8
	//ClassificationString string
	ScanAngle() int16
	UserData() uint8
	PointSourceID() uint16
	GpsTime() float64
	Red() uint16
	Green() uint16
	Blue() uint16
	NIR() uint16
	WavePacketDesc() uint8
	WaveOffset() uint64
	WaveSize() uint32
	//ReturnPointWaveLoc uint32
	X_t() float32
	Y_t() float32
	Z_t() float32
}

type Point struct {
	RecordFormat   uint8
	X              float64
	Y              float64
	Z              float64
	Intensity      uint16
	RetNum         uint8
	RetCount       uint8
	ScanChannel    uint8
	ScanDirection  uint8
	Edge           uint8
	Classification uint8
	ScanAngle      int16
	UserData       uint8
	PointSourceID  uint16
	GpsTime        float64
	Red            uint16
	Green          uint16
	Blue           uint16
	NIR            uint16
	WavePacketDesc uint8
	WaveOffset     uint64
	WaveSize       uint32
	X_t            float32
	Y_t            float32
	Z_t            float32
}

/*
All lidar point formats are based on a combination of the following structures

- Point format 0 or 6
- GPS time // mandatory in 6-10
- Red, green and blue packet
- NIR packet // 6-10 only
- Wave packet

Each point format has an internal, binary representation of the point for
reading/writing to disk (pointPacketN, where N is 0-10), and the implementation
of the format which satisfies the Pointer interface (pointFormatN).  Omitted
fields are zeroed out (not initialized), and return 0.  You must override
RecordFormat() for each concrete type.

The members are public so we can use binary.Read for reading, but the structs
are private.
*/

type pointPacket0 struct {
	X_, Y_, Z_ int32
	Intensity_ uint16
	// Return Number 0-2
	// Number of returns 3-5
	// Scan Direction flag 6
	// Edge of flight line 7
	RetScanData_    uint8
	Classification_ uint8
	ScanAngle_      int8
	UserData_       uint8
	PointSourceId_  uint16
}

type pointPacket5 struct {
	pointPacket0
	GpsTimePacket
	rgbPacket
	nirPacket
	wavePacket
}

type pointPacket6 struct {
	X_, Y_, Z_ int32
	Intensity_ uint16
	// Return Number 0-4
	// Number of returns 4-7
	ReturnData_ uint8
	// Classification Flags 0-3
	// Scanner Channel 4-5
	// Scan Direction flag 6
	// Edge of flight line 7
	ScanData_       uint8
	Classification_ uint8
	UserData_       uint8
	ScanAngle_      int16
	PointSourceId_  uint16
	GpsTimePacket
}

type GpsTimePacket float64

type rgbPacket struct {
	Red_, Green_, Blue_ uint16
}

type nirPacket uint16

type wavePacket struct {
	WavePacketDesc_      uint8
	WaveformOffset_      uint64
	WaveformSize_        uint32
	WaveformReturnPoint_ float32
	X_t_, Y_t_, Z_t_     float32
}

func (las *Lasf) point5ToPoint(p5 pointPacket5) Point {
	var p Point
	p.RecordFormat = las.PointFormat
	p.X = float64(p5.X_)*las.XScale + las.XOffset
	p.Y = float64(p5.Y_)*las.YScale + las.YOffset
	p.Z = float64(p5.Z_)*las.ZScale + las.ZOffset
	p.Intensity = p5.Intensity_
	p.RetNum = convertToUInt8(p5.RetScanData_, 0, 3)
	p.RetCount = convertToUInt8(p5.RetScanData_, 3, 3)
	// Skip scanner channel
	p.ScanDirection = convertToUInt8(p5.RetScanData_, 6, 1)
	p.Edge = convertToUInt8(p5.RetScanData_, 7, 1)
	p.Classification = p5.Classification_
	p.ScanAngle = int16(p5.ScanAngle_)
	p.UserData = p5.UserData_
	p.PointSourceID = p5.PointSourceId_
	//p.GpsTime = p5.GpsTimePacket
	p.Red = p5.Red_
	p.Green = p5.Green_
	p.Blue = p5.Blue_
	//p.NIR = p5.NIR_
	//p.WavePacketDesc = p5.WavePacketDesc_
	//p.WaveOffset = p5.WaveOffset_
	//p.WaveSize = p5.WaveSize_
	//p.X_t = p5.X_t_
	//p.Y_t = p5.Y_t_
	//p.Z_t = p5.Z_t_
	return p
}

func (las *Lasf) readPoint(i uint64) (Point, error) {
	offset := uint64(las.PointOffset) + uint64(las.PointSize)*i
	las.fin.Seek(int64(offset), os.SEEK_SET)
	var p0 pointPacket0
	var p6 pointPacket6
	var gps GpsTimePacket
	var rgb rgbPacket
	var nir nirPacket
	var wave wavePacket
	var err error
	_ = err
	f := las.PointFormat
	switch {
	case f >= 0 && f < 6:
		err = binary.Read(las.fin, binary.LittleEndian, &p0)
		if f == 1 || f == 3 || f == 4 || f == 5 {
			err = binary.Read(las.fin, binary.LittleEndian, &gps)
		}
		if f == 2 || f == 3 || f == 5 {
			err = binary.Read(las.fin, binary.LittleEndian, &rgb)
		}
		if f == 4 || f == 5 {
			err = binary.Read(las.fin, binary.LittleEndian, &wave)
		}
		p := pointPacket5{pointPacket0: p0, GpsTimePacket: gps, rgbPacket: rgb, wavePacket: wave}
		return las.point5ToPoint(p), nil
	case f >= 6 && f <= 10:
		err = binary.Read(las.fin, binary.LittleEndian, &p6)
		if f == 7 || f == 8 || f == 10 {
			err = binary.Read(las.fin, binary.LittleEndian, &rgb)
		}
		if f == 8 || f == 10 {
			err = binary.Read(las.fin, binary.LittleEndian, &nir)
		}
		if f == 9 || f == 10 {
			err = binary.Read(las.fin, binary.LittleEndian, &wave)
		}
		return Point{}, nil
	default:
		panic("Invalid point format")
	}
}


