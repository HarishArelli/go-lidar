package las

type header13 struct {
	header12
	WaveformOffset_ uint64
}

func (h *header13) WaveformOffset() uint64 {
	return h.WaveformOffset_
}
