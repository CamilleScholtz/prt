package ports

import (
	"os"
)

// An Permission is a type describing the permission bits found in the `.md5sum`
// file of a port.
type Permission struct {
	FileMode os.FileMode
	// TODO: Create a custom type for this.
	Owner string
}

func toFileMode(str string) os.FileMode {
	var v os.FileMode
	switch str[0:1] {
	case "d":
		v = os.ModeDir
	case "a":
		v = os.ModeAppend
	case "l":
		v = os.ModeExclusive
	case "T":
		v = os.ModeTemporary
	case "L":
		v = os.ModeSymlink
	case "D":
		v = os.ModeDevice
	case "p":
		v = os.ModeNamedPipe
	case "S":
		v = os.ModeSocket
	case "u":
		v = os.ModeSetuid
	case "g":
		v = os.ModeSetgid
	case "c":
		v = os.ModeCharDevice
	case "t":
		v = os.ModeSticky
	case "-":
		v = 0
	}
	v |= parseSegment(str[7:10])
	v |= (parseSegment(str[4:7]) << 3)
	v |= (parseSegment(str[1:4]) << 6)

	return os.FileMode(v)
}

func parseSegment(seg string) os.FileMode {
	var v os.FileMode

	if seg[0] == 'r' {
		v += 4
	}
	if seg[1] == 'w' {
		v += 2
	}
	if seg[2] == 'x' {
		v += 1
	}

	return v
}
