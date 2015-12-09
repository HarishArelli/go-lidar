// Copyright (c) 2015 Boise Center Aerospace Laboratory.
// All rights reserved.  // Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.


package las

type header14 struct {
	header13
	EvlrOffset_     uint64
	EvlrCount_      uint32
	PointCount_     uint64
	PointsByReturn_ [15]uint64
}

func (h *header14) PointCount() uint64 {
	return uint64(h.PointCount_)
}

func (h *header14) PointsByReturn() []uint64 {
	return []uint64(h.PointsByReturn_[:])
}

func (h *header14) EvlrOffset() uint64 {
	return h.EvlrOffset_
}

func (h *header14) EvlrCount() uint32 {
	return h.EvlrCount_
}
