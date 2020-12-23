package assets

import (
	"embed"
	"github.com/brunoga/blinks-installer/assets/avrdude"
	"github.com/brunoga/blinks-installer/assets/hex"
)

func Get() map[string]embed.FS {
	return map[string]embed.FS{
		"avrdude": avrdude.Get(),
		"hex":     hex.Get(),
	}
}
