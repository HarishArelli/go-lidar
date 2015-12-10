package pulsewaves

import (
	"testing"
)

const fname = "/home/kyle/src/bsu/lidar/data/riegl_example4.pls"

func TestOpen(t *testing.T) {
	pw, err := Open(fname)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	for i := 0; i < 10; i++ {
		_ = pw.ReadPoint(uint64(i))
	}
}
