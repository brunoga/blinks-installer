package hex

import (
	"embed"
)

//go:embed *.hex
var fs embed.FS

func Get() embed.FS {
	return fs
}
