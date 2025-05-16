package fshelp

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func FindPkgDir() (string, error) {
	// Look at our lineage to find "pkg"
	var accum string
	for i := 0; i < 7; i++ { // only go up a few subdirs
		accum = filepath.Join("..", accum)

		fds, err := os.ReadDir(accum)
		if err != nil {
			return "", fmt.Errorf("listing %s failed: %w", accum, err)
		}

		for idx := range fds {
			if fds[idx].IsDir() && fds[idx].Name() == "pkg" {
				return filepath.Join(accum, "pkg"), nil
			}
		}
	}
	return "", errors.New("no pkg ancestor found")
}
