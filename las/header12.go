// Copyright (c) 2015 Boise Center Aerospace Laboratory.
// All rights reserved.  Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package las

type header12 struct {
	Signature_ [4]byte // Must be "LASF"

	// File Source ID and Global Encoding were both available in 1.2.  They were
	// reserved at various points in time before that.  Instead of another
	// struct, ignore them if the version dictates they weren't available, then
	// we don't need header10 and header11.
	FileSourceId_                            uint16
	GlobalEncoding_                          uint16
	ProjectID1_                              uint32
	ProjectID2_                              uint16
	ProjectID3_                              uint16
	ProjectID4_                              [8]byte
	VMaj_                                    uint8
	VMin_                                    uint8
	SysIdentifier_                           [32]byte
	GenSoftware_                             [32]byte
	CreateDOY_                               uint16
	CreateYear_                              uint16
	HeaderSize_                              uint16
	PointOffset_                             uint32
	VlrCount_                                uint32
	PointFormat_                             uint8
	PointSize_                               uint16
	LegacyPointCount_                        uint32
	LegacyPointsByReturn_                    [5]uint32
	XScale_, YScale_, ZScale_                float64
	XOffset_, YOffset_, ZOffset_             float64
	MaxX_, MinX_, MaxY_, MinY_, MaxZ_, MinZ_ float64
}

func (h *header12) Signature() [4]byte {
	return h.Signature_
}

func (h *header12) FileSourceId() uint16 {
	return h.FileSourceId_
}

func (h *header12) GlobalEncoding() uint16 {
	return h.GlobalEncoding_
}

func (h *header12) ProjectID1() uint32 {
	return h.ProjectID1_
}
func (h *header12) ProjectID2() uint16 {
	return h.ProjectID2_
}

func (h *header12) ProjectID3() uint16 {
	return h.ProjectID3_
}

func (h *header12) ProjectID4() [8]byte {
	return h.ProjectID4_
}

func (h *header12) VMaj() uint8 {
	return h.VMaj_
}

func (h *header12) VMin() uint8 {
	return h.VMin_
}

func (h *header12) SysIdentifier() [32]byte {
	return h.SysIdentifier_
}

func (h *header12) GenSoftware() [32]byte {
	return h.GenSoftware_
}

func (h *header12) CreateDOY() uint16 {
	return h.CreateDOY_
}

func (h *header12) CreateYear() uint16 {
	return h.CreateYear_
}

func (h *header12) HeaderSize() uint16 {
	return h.HeaderSize_
}

func (h *header12) PointOffset() uint32 {
	return h.PointOffset_
}

func (h *header12) VlrCount() uint32 {
	return h.VlrCount_
}

func (h *header12) PointFormat() uint8 {
	return h.PointFormat_
}

func (h *header12) PointSize() uint16 {
	return h.PointSize_
}

func (h *header12) PointCount() uint64 {
	return uint64(h.LegacyPointCount_)
}

func (h *header12) PointsByReturn() []uint64 {
	p := make([]uint64, 5)
	for i, v := range h.LegacyPointsByReturn_ {
		p[i] = uint64(v)
	}
	return p
}

func (h *header12) XScale() float64 {
	return h.XScale_
}

func (h *header12) YScale() float64 {
	return h.YScale_
}

func (h *header12) ZScale() float64 {
	return h.ZScale_
}

func (h *header12) XOffset() float64 {
	return h.XOffset_
}

func (h *header12) YOffset() float64 {
	return h.YOffset_
}

func (h *header12) ZOffset() float64 {
	return h.ZOffset_
}

func (h *header12) MaxX() float64 {
	return h.MaxX_
}

func (h *header12) MinX() float64 {
	return h.MinX_
}

func (h *header12) MaxY() float64 {
	return h.MaxY_
}

func (h *header12) MinY() float64 {
	return h.MinY_
}

func (h *header12) MaxZ() float64 {
	return h.MaxZ_
}

func (h *header12) MinZ() float64 {
	return h.MinZ_
}

func (h *header12) WaveformOffset() uint64 {
	return 0
}

func (h *header12) EvlrOffset() uint64 {
	return 0
}

func (h *header12) EvlrCount() uint32 {
	return 0
}
