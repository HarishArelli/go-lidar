// Copyright (c) 2015 Boise Center Aerospace Laboratory.
// All rights reserved.  Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package las

import (
	"encoding/binary"
	"io"
	"strings"
)

type vlrHeader struct {
	_             uint16 // Reserved
	UserId_       [16]byte
	RecordId_     uint16
	RecordLength_ uint16
	Description_  [32]byte
}

func (vlr vlrHeader) UserId() [16]byte {
	return vlr.UserId_
}

func (vlr vlrHeader) RecordId() uint16 {
	return vlr.RecordId_
}

func (vlr vlrHeader) RecordLength() uint16 {
	return vlr.RecordLength_
}

func (vlr vlrHeader) Description() [32]byte {
	return vlr.Description_
}

type vlr interface {
	UserId() [16]byte
	RecordId() uint16
	RecordLength() uint16
	Description() [32]byte
}

type lookupVlr struct {
	vlr []classificationVlr
}

type classificationVlr struct {
	ClassNumber [16]byte
	Description [16]byte
}

type TextAreaDescVlr struct {
}

type extraBytesVlr struct {
	Reserved    [16]byte
	Data_type   [8]byte
	Options     [8]byte
	Name        [32]byte
	Unused      [4]byte
	Nodata      [24]byte
	Min         [24]byte
	Max         [24]byte
	Scale       [24]byte
	Offset      [24]byte
	Description [32]byte
}

type extraBytesDescriptor struct {
	extraBytes []extraBytesVlr
}

type suspersededVlr struct {
}

type waveformpacketDescVlr struct {
}

type geoKeyEntry struct {
	KeyID_      uint16
	TiffTagLoc_ uint16
	Count_      uint16
	Offset_     uint16
}

type geoKeys struct {
	KeyDirVersion_  uint16
	KeyDirRevision_ uint16
	MinorRevision_  uint16
	KeyCount_       uint16
	Keys_           []geoKeyEntry
}

// Read vlrs and store predefined vlrs in memory.
func (l *Lasf) readVlrs() error {
	vlrCount := int(l.VlrCount())
	l.fin.Seek(int64(l.HeaderSize()), 0)
	var vlrh vlrHeader
	for i := 0; i < vlrCount; i++ {
		err := binary.Read(l.fin, binary.LittleEndian, &vlrh)
		userId := string(vlrh.UserId_[:])
		recordId := vlrh.RecordId_
		if err != nil {
			return err
		}
		// Handle pre-defined vlrs
		if strings.HasPrefix(userId, "LASF_Spec") {
			switch {
			// lookup classification
			case recordId == 0:
				var lookVlr lookupVlr
				err := binary.Read(l.fin, binary.LittleEndian, &lookVlr)
				if err != nil {
					return err
				}
			// Text Area Description
			case recordId == 3:
				var TextAreaVlr TextAreaDescVlr
				err := binary.Read(l.fin, binary.LittleEndian, &TextAreaVlr)
				if err != nil {
					return err
				}
			//Extra Bytes
			case recordId == 4:
				var extra extraBytesVlr
				err := binary.Read(l.fin, binary.LittleEndian, &extra)
				if err != nil {
					return err
				}
			// Susperseded
			case recordId == 7:
				var susperse suspersededVlr
				err := binary.Read(l.fin, binary.LittleEndian, &susperse)
				if err != nil {
					return err
				}
			// WaveForm Packet Descriptor
			case recordId > 99 && recordId < 355:
				var waveformpacket waveformpacketDescVlr
				err := binary.Read(l.fin, binary.LittleEndian, &waveformpacket)
				if err != nil {
					return err
				}
				// WaveForm Data Packets
				/*
					case recordId == 65535:
						var waveformdata  waveformdataVlr
						err := binary.Read(l.fin, binary.LittleEndian, &waveformdata)
						if err != nil {
							return err
						}
				*/
			}
		} else if strings.HasPrefix(userId, "LASF_Proj") {
			switch recordId {
			case 34735:
				fallthrough
			case 34736:
				fallthrough
			case 34737:
				l.fin.Seek(int64(vlrh.RecordLength_), 1)
				continue
			case 2111:
				fallthrough
			case 2112:
				data := make([]byte, vlrh.RecordLength_)
				n, err := io.ReadAtLeast(l.fin, data, int(vlrh.RecordLength_))
				if n < 1 || err != nil {
					return err
				}
			}
		} else {
			data := make([]byte, vlrh.RecordLength_)
			n, err := io.ReadAtLeast(l.fin, data, int(vlrh.RecordLength_))
			if n < 1 || err != nil {
				return err
			}
		}
	}
	return nil
}
