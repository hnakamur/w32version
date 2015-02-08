package w32version

import (
	"errors"
	"syscall"
)

type W32Version uint16

var UnknownWindowsVersion = errors.New("Unknown Windows version")

const (
	WindowsVista = 0x0006
	Windows7     = 0x0106
	Windows8     = 0x0206
	Windows8_1   = 0x0306
	Windows10    = 0x0406
)

func GetVersion() (W32Version, error) {
	v, err := syscall.GetVersion()
	if err != nil {
		return 0, err
	}
	v = v & 0xffff
	switch v {
	case WindowsVista, Windows7, Windows8, Windows8_1, Windows10:
		return W32Version(v), nil
	default:
		return 0, UnknownWindowsVersion
	}
}

func (v W32Version) String() string {
	switch v {
	case WindowsVista:
		return "Vista"
	case Windows7:
		return "7"
	case Windows8:
		return "8"
	case Windows8_1:
		return "8.1"
	case Windows10:
		return "10"
	default:
		panic(UnknownWindowsVersion)
	}
}
