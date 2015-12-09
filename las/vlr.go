// Copyright (c) 2015 Boise Center Aerospace Laboratory.
// All rights reserved.  Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.


package las

import (
	"encoding/binary"
	"fmt"
	"os"
)

type vlrHeader struct {
	_             uint16 // Reserved
	UserId_       [16]byte
	RecordId_     uint16
	RecordLength_ uint16
	Description_  [32]byte
}

func (v *vlrHeader) Description() string {
	return string(v.Description_[:])
}

func (las *Lasf) readVlrData() error {
	offset := int64(las.HeaderSize())
	las.fin.Seek(offset, os.SEEK_SET)
	var header vlrHeader
	err := binary.Read(las.fin, binary.LittleEndian, &header)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", header)
	return nil
}
