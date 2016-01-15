package las

import (
	"fmt"
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

func getLas(i int) (*cLas, error) {
	cLasMutex.Lock()
	defer cLasMutex.Unlock()
	lf, ok := lasMap.m[i]
	if !ok {
		return nil, fmt.Errorf("Invalid File Handle")
	}
	return lf, nil
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
	lf, err := getLas(fid)
	if err != nil {
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
	lf, err := getLas(fid)
	if err != nil {
		return LASF_INVALIDHANDLE
	}
	if lf.p == nil {
		return LASF_INVALIDPOINT
	}
	*x = lf.p.X()
	return LASF_OK
}

func LasfPointY(fid int, y *float64) int {
	lf, err := getLas(fid)
	if err != nil {
		return LASF_INVALIDHANDLE
	}
	if lf.p == nil {
		return LASF_INVALIDPOINT
	}
	*y = lf.p.Y()
	return LASF_OK
}

func LasfPointZ(fid int, z *float64) int {
	lf, err := getLas(fid)
	if err != nil {
		return LASF_INVALIDHANDLE
	}
	if lf.p == nil {
		return LASF_INVALIDPOINT
	}
	*z = lf.p.Z()
	return LASF_OK
}
