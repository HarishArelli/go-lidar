package pulsewaves

import (
	"encoding/binary"
	"fmt"
	"os"
	"strings"
)

// This format is cumbersome.  There must be at least 1 avlr, the end marker
// avlr.  We are checking for that here.  We will add semantics to check for
// the rest of the avlrs later.
func (p *PulseWave) readAvlrs() error {
	var header avlrHeader
	offset := p.pHeader.PulseOffset
	offset += p.pHeader.PulseCount * uint64(p.pHeader.PulseSize)
	p.pin.Seek(int64(offset), os.SEEK_SET)
	err := binary.Read(p.pin, binary.LittleEndian, &header)
	if err != nil {
		return err
	}
	if !strings.HasPrefix(string(header.UserId[:]), "PulseWaves_Spec") || (header.RecordId <= 10000 && header.RecordId > 100255) {
		return fmt.Errorf("Could not find avlr end marker")
	}
	return nil
}
