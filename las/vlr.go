// Copyright (c) 2015 Boise Center Aerospace Laboratory.
// All rights reserved.  Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package las

type vlrHeader struct {
	_             uint16 // Reserved
	UserId_       [16]byte
	RecordId_     uint16
	RecordLength_ uint16
	Description_  [32]byte
}
