package pulsewaves

var plsSignature = [15]byte{'P', 'u', 'l', 's', 'e',
	'W', 'a', 'v', 'e', 's',
	'P', 'u', 'l', 's', 'e'}

// Header of the *.pls file.  Reserved fields are ignored with '_'
type plsHeader struct {
	Signature                          [15]byte
	_                                  byte // Should be '\0' for signature terminator
	GlobalParam                        uint32
	FileSourceID                       uint32
	ProjectGUID1                       uint32
	ProjectGUID2                       uint16
	ProjectGUID3                       uint16
	ProjectGUID4                       [8]byte
	SysIdentifier                      [64]byte
	GeneratingSoftware                 [64]byte
	FileCreateDay                      uint16
	FileCreateYear                     uint16
	VersionMajor                       uint8
	VersionMinor                       uint8
	HeaderSize                         uint16
	PulseOffset                        uint64
	PulseCount                         uint64
	PulseFormat                        uint32
	PulseAttributes                    uint32
	PulseSize                          uint32
	PulseCompression                   uint32
	_                                  int64 // Reserved
	VlrCount                           uint32
	AvlrCount                          uint32
	GpsTimeScale                       float64
	GpsTimeOffset                      float64
	GpsTimeMin                         float64
	GpsTimeMax                         float64
	XScale, YScale, ZScale             float64
	XOffset, YOffset, ZOffset          float64
	XMin, XMax, YMin, YMax, ZMin, ZMax float64
}

type vlrHeader struct {
	UserId       [16]byte
	RecordId     uint32
	_            uint32 // Reserved
	RecordLength uint64
	Description  [64]byte
}

type avlrHeader struct {
	UserId       [16]byte
	RecordId     uint32
	_            uint32 // Reserved
	RecordLength uint64
	Description  [64]byte
}

var wvsSignature = [15]byte{'P', 'u', 'l', 's', 'e',
	'W', 'a', 'v', 'e', 's',
	'W', 'a', 'v', 'e', 's'}

type wvsHeader struct {
	Signature   [15]byte
	_           byte // Should be '\0' for signature terminator
	Compression uint32
	_           [40]byte // Reserved
}
