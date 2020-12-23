package windows_amd64

import "embed"

//go:embed avrdude.exe libusb0.dll avrdude.conf
var AvrdudeFS embed.FS
