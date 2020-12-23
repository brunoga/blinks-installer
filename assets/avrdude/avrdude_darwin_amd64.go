package avrdude

import (
	"embed"
	"github.com/brunoga/blinks-installer/assets/avrdude/darwin_amd64"
)

func Get() embed.FS {
	return darwin_amd64.AvrdudeFS
}

func GetBinaryName() string {
	return "avrdude"
}
