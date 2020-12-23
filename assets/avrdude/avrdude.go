// +build !linux,!windows,!darwin !amd64

package avrdude

func Get() embed.FS {
	return nil
}
