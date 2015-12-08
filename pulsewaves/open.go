package pulsewaves

import (
	"encoding/binary"
	"fmt"
	"os"
	"path"
	"strings"
)

func findWaveFile(pfile string) string {
	folder := path.Dir(pfile)
	base := path.Base(pfile)
	base = base[:strings.Index(base, path.Ext(base))]
	wfile := path.Join(folder, base+".wvs")
	return wfile
}

func Open(filename string) (*PulseWave, error) {
	var p PulseWave
	pname := filename
	ext := path.Ext(pname)
	if ext != ".pls" {
		return nil, fmt.Errorf("%s is not a valid extension for a PulseWave file", ext)
	}
	pin, err := os.Open(pname)
	if err != nil {
		return nil, err
	}
	// Check for the waveform file
	wfile := findWaveFile(pname)
	win, err := os.Open(wfile)
	if err != nil {
		return nil, err
	}
	p.pname = pname
	p.wname = wfile
	p.pin = pin
	p.win = win

	err = p.readHeader()
	if err != nil {
		return nil, err;
	}
	err = p.readVlrs()
	if err != nil {
		return nil, err;
	}

	return &p, nil
}

func (p *PulseWave) Close() error {
	//p.pin.Close()
	//p.win.Close()
	p.pHeader = nil
	p.wHeader = nil
	p.p = nil
	p.index = 0
	return nil
}

func (p *PulseWave) readHeader() error {

	var ph plsHeader
	var wh wvsHeader

	if p.pin == nil {
		return fmt.Errorf("Invalid pulse wave file: %s", p.pname)
	}
	if p.win == nil {
		return fmt.Errorf("Invalid wave file: %s", p.wname)
	}
	err := binary.Read(p.pin, binary.LittleEndian, &ph)
	if err != nil {
		return err
	}
	if ph.Signature != plsSignature {
		return fmt.Errorf("Invalid signature in PulseWave file: %s", ph.Signature)
	}
	err = binary.Read(p.win, binary.LittleEndian, &wh)
	if err != nil {
		return err
	}
	if wh.Signature != wvsSignature {
		return fmt.Errorf("Invalid signature in Wave file: %s", wh.Signature)
	}

	p.pHeader = &ph
	p.wHeader = &wh
	return nil
}
