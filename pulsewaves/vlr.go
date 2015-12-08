package pulsewaves

import (
	"encoding/binary"
	"io"
	"strings"
)

// Types and methods for various VLRs.  Not that each VLR has a 'public' struct
// for exact storage representation, and may consist of one or more 'private'
// structs used with reflection and binary reading.  The latter have 'Vlr'
// appended to the name.
type scannerVlr struct {
	Size             uint32
	_                uint32
	Instrument       [64]byte
	Serial           [64]byte
	WaveLength       float32
	OutPulseWidth    float32
	ScanPattern      uint32
	MirrorFacetCount uint32
	ScanFreq         float32
	ScanAngleMin     float32
	ScanAngleMax     float32
	PulseFreq        float32
	BeamDiamExit     float32
	BeamDivergence   float32
	MinRange         float32
	MaxRange         float32
	Description      [64]byte
}

type scanner struct {
	vlr scannerVlr
}

type pulseDescVlr struct {
	WaveSize       uint32
	_              uint32 // Reserved
	OpticalCenter  int32
	ExtraWaveBytes uint16
	SampleCount    uint16
	SampleUnits    float32 // nanoseconds
	Compression    uint32
	ScannerIndex   uint32
	Description    [64]byte
}

type sampleRecordVlr struct {
	Size            uint32
	_               uint32
	Type            uint8
	Channel         uint8
	_               uint8
	BitsDurAnchor   uint8
	ScaleDurAnchor  float32
	OffsetDurAnchor float32
	BitsNumSegments uint8
	BitsNumSamples  uint8
	SegmentCount    uint16
	SampleCount     uint32
	BitsSample      uint16
	LookupIndex     uint16
	SampleUnits     float32
	Compression     uint32
	Description     [64]byte
}

type pulseDescriptor struct {
	vlr           pulseDescVlr
	sampleRecords []sampleRecordVlr
}

type tableRecordVlr struct {
	Size        uint32
	_           uint32 // Reserved
	TableCount  uint32
	Description [64]byte
}

type lookupTableRecordVlr struct {
	Size           uint32
	_              uint32 // Reserved
	EntryCount     uint32
	MeasurmentUnit uint16
	DataType       uint8
	Options        uint8
	Compression    uint32
	Description    [64]byte
}

type geoKeyEntry struct {
	KeyID      uint16
	TiffTagLoc uint16
	Count      uint16
	Offset     uint16
}

type geoKeys struct {
	KeyDirVersion  uint16
	KeyDirRevision uint16
	MinorRevision  uint16
	KeyCount       uint16
	Keys           []geoKeyEntry
}

// Read vlrs and store predefined vlrs in memory.
func (p *PulseWave) readVlrs() error {
	vlrCount := int(p.pHeader.VlrCount)
	p.pin.Seek(int64(p.pHeader.HeaderSize), 0)
	var vlrh vlrHeader
	for i := 0; i < vlrCount; i++ {
		err := binary.Read(p.pin, binary.LittleEndian, &vlrh)
		userId := string(vlrh.UserId[:])
		recordId := vlrh.RecordId
		if err != nil {
			return err
		}
		// Handle pre-defined vlrs
		if strings.HasPrefix(userId, "PulseWaves_Spec") {
			switch {
			case recordId > 100000 && recordId < 100255:
				var scanVlr scannerVlr
				err := binary.Read(p.pin, binary.LittleEndian, &scanVlr)
				if err != nil {
					return err
				}
			// Pulse descriptor
			case recordId > 200000 && recordId < 200255:
				var pulseVlr pulseDescVlr
				var pulseDesc pulseDescriptor
				err := binary.Read(p.pin, binary.LittleEndian, &pulseVlr)
				if err != nil {
					return err
				}
				var sample sampleRecordVlr
				for i := 0; i < int(pulseVlr.SampleCount); i++ {
					err := binary.Read(p.pin, binary.LittleEndian, &sample)
					if err != nil {
						return err
					}
					pulseDesc.sampleRecords = append(pulseDesc.sampleRecords, sample)
				}
				p.pulseDescriptors = append(p.pulseDescriptors, pulseDesc)
			case recordId > 300000 && recordId < 300255:
				var table tableRecordVlr
				err := binary.Read(p.pin, binary.LittleEndian, &table)
				if err != nil {
					return err
				}
				var lookup lookupTableRecordVlr
				for i := 0; i < int(table.TableCount); i++ {
					err := binary.Read(p.pin, binary.LittleEndian, &lookup)
					if err != nil {
						return err
					}
				}
			}
		} else if strings.HasPrefix(userId, "PulseWaves_Proj") {
			switch recordId {
			case 34735:
				fallthrough
			case 34736:
				fallthrough
			case 34737:
				p.pin.Seek(int64(vlrh.RecordLength), 1)
				continue
			case 2111:
				fallthrough
			case 2112:
				data := make([]byte, vlrh.RecordLength)
				n, err := io.ReadAtLeast(p.pin, data, int(vlrh.RecordLength))
				if n < 1 || err != nil {
					return err
				}
			}
		} else {
			data := make([]byte, vlrh.RecordLength)
			n, err := io.ReadAtLeast(p.pin, data, int(vlrh.RecordLength))
			if n < 1 || err != nil {
				return err
			}
		}
	}
	return nil
}
