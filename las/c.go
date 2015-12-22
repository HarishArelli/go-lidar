package las

import (
	"sync"
)

var cLasMutex sync.Mutex

type cLas struct {
	handle int
	fname  string
	//references int
	lasf *Lasf
	p    Pointer
}

type lasFileMap struct {
	m map[int]*cLas
	i int
}

var lasMap lasFileMap

func init() {
	lasMap.m = make(map[int]*cLas)
}

const (
	LASF_OK = iota
	LASF_CANTOPEN
	LASF_CANTREAD
	LASF_INVALIDHANDLE
	LASF_INVALIDINDEX
	LASF_INVALIDPOINT
	LASF_ERROR
)

func LasfOpen(fname string, fid *int) int {
	cLasMutex.Lock()
	defer cLasMutex.Unlock()
	l, err := Open(fname)
	if err != nil {
		return LASF_CANTOPEN
	}
	c := new(cLas)
	c.handle = lasMap.i
	c.fname = fname
	c.lasf = l
	lasMap.m[lasMap.i] = c
	*fid = lasMap.i
	lasMap.i++
	return LASF_OK
}

func LasfReadNextPoint(fid int) int {
	cLasMutex.Lock()
	lf, ok := lasMap.m[fid]
	cLasMutex.Unlock()
	if !ok {
		return LASF_INVALIDHANDLE
	}
	p, e := lf.lasf.GetNextPoint()
	if e != nil {
		return LASF_INVALIDINDEX
	}
	lf.p = p
	return LASF_OK
}

func LasfPointX(fid int, x *float64) int {
	cLasMutex.Lock()
	lf, ok := lasMap.m[fid]
	cLasMutex.Unlock()
	if !ok {
		return LASF_INVALIDHANDLE
	}
	if lf.p == nil {
		return LASF_INVALIDPOINT
	}
	px := lf.p.X()
	*x = px
	return LASF_OK
}

func LasfPointY(fid int, y *float64) int {
	cLasMutex.Lock()
	lf, ok := lasMap.m[fid]
	cLasMutex.Unlock()
	if !ok {
		return LASF_INVALIDHANDLE
	}
	if lf.p == nil {
		return LASF_INVALIDPOINT
	}
	py := lf.p.Y()
	*y = py
	return LASF_OK
}
