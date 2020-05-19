// +build !linux

package flog

import (
	"os"
)

func chown(_ string, _ os.FileInfo) error {
	return nil
}
