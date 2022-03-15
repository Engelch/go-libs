package go_libs

import (
	"path/filepath"
)

func FilenameWithoutSuffix(filename string) string {
	extension := filepath.Ext(filename)
	return filename[0 : len(filename)-len(extension)]
}

// eof
