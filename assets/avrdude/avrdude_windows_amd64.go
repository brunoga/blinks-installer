package avrdude

import (
	"embed"
	"github.com/brunoga/blinks-installer/assets/avrdude/windows_amd64"
)

func Get() embed.FS {
	return windows_amd64.AvrdudeFS
}

func GetBinaryName() string {
	return "avrdude.exe"
}
