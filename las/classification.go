// Copyright (c) 2015 Boise Center Aerospace Laboratory.
// All rights reserved.  // Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.


package las

// String representations of classification values

var classification05 = [...]string{"Created, never classified",
	"Unclassified",
	"Low Vegetation",
	"Medium Vegetation",
	"High Vegetation",
	"Building",
	"Low Point (noise)",
	"Model Key-point (mass point)",
	"Water",
	"Reserved for ASPRS Definition",
	"Reserved for ASPRS Definition",
	"Overlap Point",
	"for ASPRS Definition"}

var classification610 = [...]string{"Created, never classified",
	"Unclassified",
	"Ground",
	"Low Vegetation",
	"Medium Vegetation",
	"High Vegetation",
	"Building",
	"Low Point (noise)",
	"Reserved",
	"Water",
	"Rail",
	"Road Surface",
	"Reserved ",
	"Wire – Guard (Shield)",
	"Wire – Conductor (Phase)",
	"Transmission Tower",
	"Wire-structure Connector (e.g. Insulator)",
	"Bridge Deck",
	"High Noise 19-63",
	"Reserved",
	"User definable"}
