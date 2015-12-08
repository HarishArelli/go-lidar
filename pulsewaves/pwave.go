package pulsewaves

import (
	"encoding/binary"
	"fmt"
	"io"
)

type PulseWave struct {
	// Names of the pulsewave file and associated waveform file
	pname, wname string
	pin          io.ReadSeeker
	win          io.ReadSeeker

	// Header information
	pHeader *plsHeader
	wHeader *wvsHeader
	// Current index
	index uint64
	// Various vlr information
	pulseDescriptors []pulseDescriptor


	p *pulse
}

// Public API for the PulseWaves format

func (p *PulseWave) PulseFile() string {
	return p.pname
}

func (p *PulseWave) WaveFile() string {
	return p.wname
}

func (p *PulseWave) RecordCount() uint64 {
	return p.pHeader.PulseCount
}

func (p *PulseWave) ReadPoint(i uint64) error {
	if i > p.pHeader.PulseCount {
		return fmt.Errorf("Invalid pulse index: %d", i)
	}
	offset := p.pHeader.PulseOffset + i*uint64(p.pHeader.PulseSize)
	p.pin.Seek(int64(offset), 0)
	var pu pulse
	err := binary.Read(p.pin, binary.LittleEndian, &pu)
	if err != nil {
		fmt.Println(err)
		return err
	}
	//p.win.Seek(int64(pu.WaveOffset), 0)

	fmt.Printf("%#v\n", pu)
	fmt.Printf("%+v\n", p.pulseDescriptors[pu.DescriptorIndex])
	return nil
}


