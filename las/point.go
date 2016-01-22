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
type Point struct {
	RecordFormat        uint8
	X                   float64
	Y                   float64
	Z                   float64
	Intensity           uint16
	RetNum              uint8
	RetCount            uint8
	ClassificationFlags uint8
	ScanChannel         uint8
	ScanDirection       uint8
	Edge                uint8
	Classification      uint8
	ScanAngle           int16
	UserData            uint8
	PointSourceID       uint16
	GpsTime             float64
	Red                 uint16
	Green               uint16
	Blue                uint16
	NIR                 uint16
	WavePacketDesc      uint8
	WaveOffset          uint64
	WaveSize            uint32
	WaveReturnPoint     float32
	Xt                  float32
	Yt                  float32
	Zt                  float32
	las                 *Lasf
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
	X, Y, Z   int32
	Intensity uint16
	// Return Number 0-2
	// Number of returns 3-5
	// Scan Direction flag 6
	// Edge of flight line 7
	RetScanData    uint8
	Classification uint8
	ScanAngle      int8
	UserData       uint8
	PointSourceId  uint16
}

type rgbPacket struct {
	Red, Green, Blue uint16
}

type wavePacket struct {
	WavePacketDesc  uint8
	WaveOffset      uint64
	WaveSize        uint32
	WaveReturnPoint float32
	Xt, Yt, Zt      float32
}

func (las *Lasf) point5ToPoint(p0 pointPacket0, gps float64, rgb rgbPacket, wave wavePacket) Point {
	var p Point
	p.RecordFormat = las.PointFormat
	p.X = float64(p0.X)*las.XScale + las.XOffset
	p.Y = float64(p0.Y)*las.YScale + las.YOffset
	p.Z = float64(p0.Z)*las.ZScale + las.ZOffset
	p.Intensity = p0.Intensity
	p.RetNum = convertToUInt8(p0.RetScanData, 0, 3)
	p.RetCount = convertToUInt8(p0.RetScanData, 3, 3)
	// Skip classification flags
	// Skip scanner channel
	p.ScanDirection = convertToUInt8(p0.RetScanData, 6, 1)
	p.Edge = convertToUInt8(p0.RetScanData, 7, 1)
	p.Classification = p0.Classification
	p.ScanAngle = int16(p0.ScanAngle)
	p.UserData = p0.UserData
	p.PointSourceID = p0.PointSourceId
	p.GpsTime = gps
	p.Red = rgb.Red
	p.Green = rgb.Green
	p.Blue = rgb.Blue
	// Skip NIR
	p.WavePacketDesc = wave.WavePacketDesc
	p.WaveOffset = wave.WaveOffset
	p.WaveSize = wave.WaveSize
	p.WaveReturnPoint = wave.WaveReturnPoint
	p.Xt = wave.Xt
	p.Yt = wave.Yt
	p.Zt = wave.Zt
	p.las = las
	return p
}

type pointPacket6 struct {
	X, Y, Z   int32
	Intensity uint16
	// Return Number 0-4
	// Number of returns 4-7
	ReturnData uint8
	// Classification Flags 0-3
	// Scanner Channel 4-5
	// Scan Direction flag 6
	// Edge of flight line 7
	ScanData       uint8
	Classification uint8
	UserData       uint8
	ScanAngle      int16
	PointSourceId  uint16
	GpsTime        float64
}

func (las *Lasf) point10ToPoint(p6 pointPacket6, rgb rgbPacket, nir uint16, wave wavePacket) Point {
	var p Point
	p.RecordFormat = las.PointFormat
	p.X = float64(p6.X)*las.XScale + las.XOffset
	p.Y = float64(p6.Y)*las.YScale + las.YOffset
	p.Z = float64(p6.Z)*las.ZScale + las.ZOffset
	p.Intensity = p6.Intensity
	p.RetNum = convertToUInt8(p6.ReturnData, 0, 4)
	p.RetCount = convertToUInt8(p6.ReturnData, 4, 4)
	p.ClassificationFlags = convertToUInt8(p6.ScanData, 0, 4)
	p.ScanChannel = convertToUInt8(p6.ScanData, 4, 2)
	p.ScanDirection = convertToUInt8(p6.ScanData, 6, 1)
	p.Edge = convertToUInt8(p6.ScanData, 7, 1)
	p.Classification = p6.Classification
	p.ScanAngle = p6.ScanAngle
	p.UserData = p6.UserData
	p.PointSourceID = p6.PointSourceId
	p.GpsTime = p6.GpsTime
	p.Red = rgb.Red
	p.Green = rgb.Green
	p.Blue = rgb.Blue
	p.NIR = nir
	p.WavePacketDesc = wave.WavePacketDesc
	p.WaveOffset = wave.WaveOffset
	p.WaveSize = wave.WaveSize
	p.WaveReturnPoint = wave.WaveReturnPoint
	p.Xt = wave.Xt
	p.Yt = wave.Yt
	p.Zt = wave.Zt
	p.las = las
	return p
}

func (las *Lasf) readPoint(i uint64) (Point, error) {
	offset := uint64(las.PointOffset) + uint64(las.PointSize)*i
	las.fin.Seek(int64(offset), os.SEEK_SET)
	var p0 pointPacket0
	var p6 pointPacket6
	var gps float64
	var rgb rgbPacket
	var nir uint16
	var wave wavePacket
	var err error
	var p Point
	f := las.PointFormat
	switch {
	case f >= 0 && f < 6:
		err = binary.Read(las.fin, binary.LittleEndian, &p0)
		if err != nil {
			return p, err
		}
		if f == 1 || f == 3 || f == 4 || f == 5 {
			err = binary.Read(las.fin, binary.LittleEndian, &gps)
			if err != nil {
				return p, err
			}
		}
		if f == 2 || f == 3 || f == 5 {
			err = binary.Read(las.fin, binary.LittleEndian, &rgb)
			if err != nil {
				return p, err
			}
		}
		if f == 4 || f == 5 {
			err = binary.Read(las.fin, binary.LittleEndian, &wave)
			if err != nil {
				return p, err
			}
		}
		p = las.point5ToPoint(p0, gps, rgb, wave)
		return p, nil
	case f >= 6 && f <= 10:
		err = binary.Read(las.fin, binary.LittleEndian, &p6)
		if err != nil {
			return p, err
		}
		if f == 7 || f == 8 || f == 10 {
			err = binary.Read(las.fin, binary.LittleEndian, &rgb)
			if err != nil {
				return p, err
			}
		}
		if f == 8 || f == 10 {
			err = binary.Read(las.fin, binary.LittleEndian, &nir)
			if err != nil {
				return p, err
			}
		}
		if f == 9 || f == 10 {
			err = binary.Read(las.fin, binary.LittleEndian, &wave)
			if err != nil {
				return p, err
			}
		}
		p = las.point10ToPoint(p6, rgb, nir, wave)
		return p, nil
	default:
		panic("Invalid point format")
	}
}
