package darwin_amd64

import "embed"

//go:embed avrdude avrdude.conf
var AvrdudeFS embed.FS
