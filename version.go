package w32version

import (
	"errors"
	"os"
	"os/exec"
	"strings"
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

const versionPrefix = "[Version "
const versionSuffix = "]"

func GetVersion() (W32Version, error) {
	cmd := os.Getenv("ComSpec")
	out, err := exec.Command(cmd, "/c", "ver").Output()
	if err != nil {
		return 0, err
	}
	outStr := string(out)
	start := strings.Index(outStr, versionPrefix)
	if start == -1 {
		return 0, UnknownWindowsVersion
	}
	outStr = outStr[start+len(versionPrefix):]
	end := strings.Index(outStr, versionSuffix)
	if end == -1 {
		return 0, UnknownWindowsVersion
	}
	s := strings.Split(outStr[:end], ".")
	switch {
	case s[0] == "6" && s[1] == "0":
		return WindowsVista, nil
	case s[0] == "6" && s[1] == "1":
		return Windows7, nil
	case s[0] == "6" && s[1] == "2":
		return Windows8, nil
	case s[0] == "6" && s[1] == "3":
		return Windows8_1, nil
	case s[0] == "6" && s[1] == "4":
		return Windows10, nil
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
