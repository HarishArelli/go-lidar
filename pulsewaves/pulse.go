package pulsewaves

import (
	"fmt"
)

type pulse struct {
	GpsTime                   int64
	WaveOffset                int64
	XAnchor, YAnchor, ZAnchor int32
	XTarget, YTarget, ZTarget int32
	FirstReturn, LastReturn   int16
	DescriptorIndex           uint8
	ScanData                  uint8
	Intensity                 uint8
	Classification            uint8
}

func (p pulse) GoString() string {
	s := fmt.Sprintf("pulse: %p\n", &p)
	s += fmt.Sprintf("\tGpsTime: %d\n", p.GpsTime)
	s += fmt.Sprintf("\tWaveOffset: %d\n", p.WaveOffset)
	s += fmt.Sprintf("\tXAnchor: %d\n", p.XAnchor)
	s += fmt.Sprintf("\tYAnchor: %d\n", p.YAnchor)
	s += fmt.Sprintf("\tZAnchor: %d\n", p.ZAnchor)
	s += fmt.Sprintf("\tXTarget: %d\n", p.XTarget)
	s += fmt.Sprintf("\tYTarget: %d\n", p.YTarget)
	s += fmt.Sprintf("\tZTarget: %d\n", p.ZTarget)
	s += fmt.Sprintf("\tFirstReturn: %d\n", p.FirstReturn)
	s += fmt.Sprintf("\tLastReturn: %d\n", p.LastReturn)
	s += fmt.Sprintf("\tDescriptorIndex: %d\n", p.DescriptorIndex)
	s += fmt.Sprintf("\tScanData: %d\n", p.ScanData)
	s += fmt.Sprintf("\tIntensity: %d\n", p.Intensity)
	s += fmt.Sprintf("\tClassification: %d\n", p.Classification)
	return s
}
