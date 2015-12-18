
//+build ignore

package main

import (
	"C"
	"github.com/bcal-lidar/go-lidar/las"
)

//export lasf_open
func lasf_open(fname string, fid *int) int {
	return las.LasfOpen(fname, fid)
}

//export lasf_read_point
func lasf_read_point(fid int) int {
	return las.LasfReadNextPoint(fid)
}

//export lasf_point_x
func lasf_point_x(fid int, x *float64) int {
	return las.LasfPointX(fid, x)
}

func main() {
}
