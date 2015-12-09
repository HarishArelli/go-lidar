// Copyright (c) 2015 Boise Center Aerospace Laboratory.
// All rights reserved.  // Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.


package las

import (
	"encoding/binary"
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
	//ClassificationString() string
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
	//ReturnPointWaveLoc() uint32
	X_t() float32
	Y_t() float32
	Z_t() float32
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
*/

// Base types
type gpsTimePacket float64

func (gps gpsTimePacket) GpsTime() float64 {
	return float64(gps)
}

func (gps *gpsTimePacket) SetGpsTime(t float64) {
	*gps = gpsTimePacket(t)
}

type rgbPacket struct {
	Red_, Green_, Blue_ uint16
}

func (rgb *rgbPacket) Red() uint16 {
	return rgb.Red_
}

func (rgb *rgbPacket) Green() uint16 {
	return rgb.Green_
}

func (rgb *rgbPacket) Blue() uint16 {
	return rgb.Blue_
}

func (rgb *rgbPacket) SetRGB(r, g, b uint16) {
	rgb.Red_, rgb.Green_, rgb.Blue_ = r, g, b
}

type nirPacket uint16

func (n nirPacket) NIR() uint16 {
	return uint16(n)
}

func (nir *nirPacket) SetNIR(n uint16) {
	*nir = nirPacket(n)
}

// Encapsulate the waveform packet description
type wavePacket struct {
	WavePacketDesc_      uint8
	WaveformOffset_      uint64
	WaveformSize_        uint32
	WaveformReturnPoint_ float32
	X_t_, Y_t_, Z_t_     float32
}

func (w *wavePacket) WavePacketDesc() uint8 {
	return w.WavePacketDesc_
}

func (w *wavePacket) WaveOffset() uint64 {
	return w.WaveformOffset_
}

func (w *wavePacket) WaveSize() uint32 {
	return w.WaveformSize_
}

func (w *wavePacket) WaveformReturnPoint() float32 {
	return w.WaveformReturnPoint_
}

func (w *wavePacket) X_t() float32 {
	return w.X_t_
}

func (w *wavePacket) Y_t() float32 {
	return w.Y_t_
}

func (w *wavePacket) Z_t() float32 {
	return w.Z_t_
}

type pointPacket0 struct {
	X_, Y_, Z_ uint32
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

func (p *pointPacket0) RecordFormat() uint8 {
	return 0
}

func (p *pointPacket0) X() float64 {
	return float64(p.X_)
}

func (p *pointPacket0) Y() float64 {
	return float64(p.Y_)
}

func (p *pointPacket0) Z() float64 {
	return float64(p.Z_)
}

func (p *pointPacket0) Intensity() uint16 {
	return p.Intensity_
}

func (p *pointPacket0) RetNum() uint8 {
	return convertToUInt8(p.RetScanData_, 0, 3)
}

func (p *pointPacket0) RetCount() uint8 {
	return convertToUInt8(p.RetScanData_, 3, 3)
}

func (p *pointPacket0) ScanFlag() uint8 {
	return convertToUInt8(p.RetScanData_, 6, 1)
}

func (p *pointPacket0) Edge() uint8 {
	return convertToUInt8(p.RetScanData_, 7, 1)
}

func (p *pointPacket0) Classification() uint8 {
	return p.Classification_
}

func (p *pointPacket0) ScanAngle() int16 {
	return int16(p.ScanAngle_)
}

func (p *pointPacket0) UserData() uint8 {
	return p.UserData_
}
func (p *pointPacket0) PointSourceID() uint16 {
	return p.PointSourceId_
}

type pointFormat0 struct {
	// Concrete
	pointPacket0
	// Zero fields
	gpsTimePacket
	rgbPacket
	nirPacket
	wavePacket
}

func (p *pointFormat0) RecordFormat() uint8 {
	return 0
}

type pointPacket1 struct {
	pointPacket0
	gpsTimePacket
}

type pointFormat1 struct {
	pointPacket1
	rgbPacket
	nirPacket
	wavePacket
}

func (p *pointFormat1) RecordFormat() uint8 {
	return 1
}

type pointPacket2 struct {
	pointPacket0
	rgbPacket
}

type pointFormat2 struct {
	pointPacket2
	gpsTimePacket
	nirPacket
	wavePacket
}

func (p *pointFormat2) RecordFormat() uint8 {
	return 2
}

type pointPacket3 struct {
	pointPacket0
	gpsTimePacket
	rgbPacket
}

type pointFormat3 struct {
	pointPacket3
	nirPacket
	wavePacket
}

func (p *pointFormat3) RecordFormat() uint8 {
	return 3
}

type pointPacket4 struct {
	pointPacket1
	wavePacket
}

type pointFormat4 struct {
	pointPacket4
	rgbPacket
	nirPacket
}

func (p *pointFormat4) RecordFormat() uint8 {
	return 4
}

type pointPacket5 struct {
	pointPacket3
	wavePacket
}

type pointFormat5 struct {
	pointPacket5
	nirPacket
}

func (p *pointFormat5) RecordFormat() uint8 {
	return 5
}

func (las *Lasf) readPointFormat0() (Pointer, error) {
	var p pointPacket0
	err := binary.Read(las.fin, binary.LittleEndian, &p)
	if err != nil {
		return nil, err
	}
	return &pointFormat0{pointPacket0: p}, nil
}

func (las *Lasf) readPointFormat1() (Pointer, error) {
	var p pointPacket1
	err := binary.Read(las.fin, binary.LittleEndian, &p)
	if err != nil {
		return nil, err
	}
	return &pointFormat1{pointPacket1: p}, nil
}

func (las *Lasf) readPointFormat2() (Pointer, error) {
	var p pointPacket2
	err := binary.Read(las.fin, binary.LittleEndian, &p)
	if err != nil {
		return nil, err
	}
	return &pointFormat2{pointPacket2: p}, nil
}

func (las *Lasf) readPointFormat3() (Pointer, error) {
	var p pointPacket3
	err := binary.Read(las.fin, binary.LittleEndian, &p)
	if err != nil {
		return nil, err
	}
	return &pointFormat3{pointPacket3: p}, nil
}

func (las *Lasf) readPointFormat4() (Pointer, error) {
	var p pointPacket4
	err := binary.Read(las.fin, binary.LittleEndian, &p)
	if err != nil {
		return nil, err
	}
	return &pointFormat4{pointPacket4: p}, nil
}

func (las *Lasf) readPointFormat5() (Pointer, error) {
	var p pointPacket5
	err := binary.Read(las.fin, binary.LittleEndian, &p)
	if err != nil {
		return nil, err
	}
	return &pointFormat5{pointPacket5: p}, nil
}

type pointPacket6 struct {
	X_, Y_, Z_ uint32
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
	gpsTimePacket
}

func (p *pointPacket6) RecordFormat() uint8 {
	return 6
}

func (p *pointPacket6) X() float64 {
	return float64(p.X_)
}

func (p *pointPacket6) Y() float64 {
	return float64(p.Y_)
}

func (p *pointPacket6) Z() float64 {
	return float64(p.Z_)
}

func (p *pointPacket6) Intensity() uint16 {
	return p.Intensity_
}

func (p *pointPacket6) RetNum() uint8 {
	return convertToUInt8(p.ReturnData_, 0, 4)
}

func (p *pointPacket6) RetCount() uint8 {
	return convertToUInt8(p.ReturnData_, 4, 4)
}

/*
func (p *pointPacket6) ClassificationFlags() uint8 {
	return convertToUInt8(p.ScanData_, 0, 4)
}
func (p *pointPacket6) ScanChannel() uint8 {
	return convertToUInt8(p.ScanData_, 4, 2)
}
*/

func (p *pointPacket6) ScanFlag() uint8 {
	return convertToUInt8(p.ScanData_, 6, 1)
}

func (p *pointPacket6) Edge() uint8 {
	return convertToUInt8(p.ReturnData_, 7, 1)
}

func (p *pointPacket6) Classification() uint8 {
	return p.Classification_
}

func (p *pointPacket6) UserData() uint8 {
	return p.UserData_
}

func (p *pointPacket6) ScanAngle() int16 {
	return p.ScanAngle_
}

func (p *pointPacket6) PointSourceID() uint16 {
	return p.PointSourceId_
}

type pointFormat6 struct {
	pointPacket6
	rgbPacket
	nirPacket
	wavePacket
}

func (p *pointFormat6) RecordFormat() uint8 {
	return 6
}

type pointPacket7 struct {
	pointPacket6
	rgbPacket
}

type pointFormat7 struct {
	pointPacket7
	nirPacket
	wavePacket
}

func (p *pointFormat7) RecordFormat() uint8 {
	return 7
}

type pointPacket8 struct {
	pointPacket7
	nirPacket
}

type pointFormat8 struct {
	pointPacket8
	wavePacket
}

func (p *pointFormat8) RecordFormat() uint8 {
	return 8
}

type pointPacket9 struct {
	pointPacket6
	wavePacket
}

type pointFormat9 struct {
	pointPacket9
	rgbPacket
	nirPacket
}

func (p *pointFormat9) RecordFormat() uint8 {
	return 9
}

type pointPacket10 struct {
	pointPacket6
	rgbPacket
	nirPacket
	wavePacket
}

type pointFormat10 struct {
	pointPacket10
}

func (p *pointFormat10) RecordFormat() uint8 {
	return 10
}

func (las *Lasf) readPointFormat6() (Pointer, error) {
	var p pointPacket6
	err := binary.Read(las.fin, binary.LittleEndian, &p)
	if err != nil {
		return nil, err
	}
	return &pointFormat6{pointPacket6: p}, nil
}

func (las *Lasf) readPointFormat7() (Pointer, error) {
	var p pointPacket7
	err := binary.Read(las.fin, binary.LittleEndian, &p)
	if err != nil {
		return nil, err
	}
	return &pointFormat7{pointPacket7: p}, nil
}

func (las *Lasf) readPointFormat8() (Pointer, error) {
	var p pointPacket8
	err := binary.Read(las.fin, binary.LittleEndian, &p)
	if err != nil {
		return nil, err
	}
	return &pointFormat8{pointPacket8: p}, nil
}

func (las *Lasf) readPointFormat9() (Pointer, error) {
	var p pointPacket9
	err := binary.Read(las.fin, binary.LittleEndian, &p)
	if err != nil {
		return nil, err
	}
	return &pointFormat9{pointPacket9: p}, nil
}

func (las *Lasf) readPointFormat10() (Pointer, error) {
	var p pointPacket10
	err := binary.Read(las.fin, binary.LittleEndian, &p)
	if err != nil {
		return nil, err
	}
	return &pointFormat10{pointPacket10: p}, nil
}
