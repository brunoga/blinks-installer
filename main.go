package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/brunoga/blinks-installer/assets"
	"github.com/brunoga/blinks-installer/assets/avrdude"
)

func extractAvrdude(tmpDir string) (string, string, error) {
	fmt.Printf("Extracting avrdude   ... ")

	avrdudeAssets := assets.Get()["avrdude"]

	err := fs.WalkDir(avrdudeAssets, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			data, err := avrdudeAssets.ReadFile(path)
			if err != nil {
				return err
			}

			fileName := filepath.Base(path)

			if fileName == avrdude.GetBinaryName() {
				err = os.WriteFile(filepath.Join(tmpDir, fileName), data, 0700)
			} else {
				err = os.WriteFile(filepath.Join(tmpDir, fileName), data, 0600)
			}

			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error!\n\n")
	} else {
		fmt.Println("Done.")
	}

	return filepath.Join(tmpDir, avrdude.GetBinaryName()), filepath.Join(
		tmpDir, "avrdude.conf"), err
}

func extractHex(tmpDir string) ([]string, error) {
	fmt.Printf("Extracting hex files ... ")

	hexAssets := assets.Get()["hex"]

	var hexPaths []string

	err := fs.WalkDir(hexAssets, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			data, err := hexAssets.ReadFile(path)
			if err != nil {
				return err
			}

			hexPath := filepath.Join(tmpDir, filepath.Base(path))

			err = os.WriteFile(hexPath, data, 0600)

			if err != nil {
				return err
			}

			hexPaths = append(hexPaths, hexPath)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error!\n\n")
	} else {
		fmt.Println("Done")
	}

	return hexPaths, err
}

func installHex(avrdudeBinaryPath, avrdudeConfigPath, hexPath string) error {
	var reply string
	var dummy string

L:
	for {
		fmt.Printf("Install %q ([Y]es/[N]o/[Q]uit)? ", strings.TrimSuffix(
			filepath.Base(hexPath), filepath.Ext(hexPath)))

		n, _ := fmt.Scanln(&reply, &dummy)
		if n == 1 {
			reply = strings.ToLower(reply)
			switch {
			case reply == "y" || reply == "yes":
				break L
			case reply == "n" || reply == "no":
				return nil
			case reply == "q" || reply == "quit":
				return fmt.Errorf("aborted")
			}
		}
	}

	fmt.Printf("Installing ... ")

	cmd := exec.Command(avrdudeBinaryPath, "-B", "5", "-v", "-patmega168pb",
		"-C", avrdudeConfigPath, "-cusbtiny", fmt.Sprintf("-Uflash:w:%s:i",
			hexPath), "-u")

	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error!\n\n")

		return fmt.Errorf("%s : %s", err, stdoutStderr)
	}

	fmt.Printf("Done.\n\n")

	return nil
}

func main() {
	fmt.Printf("Blinks Installer v1.0.0\n\n")

	tmpDir, err := os.MkdirTemp("", "blinks-installer-*")
	if err != nil {
		panic(err)
	}

	defer os.RemoveAll(tmpDir)

	avrdudeBinaryPath, avrdudeConfigPath, err := extractAvrdude(tmpDir)
	if err != nil {
		panic(err)
	}

	hexPaths, err := extractHex(tmpDir)
	if err != nil {
		panic(err)
	}

	fmt.Println("")

	for _, hexPath := range hexPaths {
		err = installHex(avrdudeBinaryPath, avrdudeConfigPath, hexPath)
		if err != nil {
			panic(err)
		}
	}
}
