package go_libs

import (
	"errors"
	"os"
	"path/filepath"
)

func FilenameWithoutSuffix(filename string) string {
	extension := filepath.Ext(filename)
	return filename[0 : len(filename)-len(extension)]
}

// WriteByteArray writes a byte array into a file. If required it does so in multiple steps.
// If all succeeds then nil is returned, otherwise an error.
func WriteByteArray(file *os.File, bytes []byte) error {
	nsum := 0
	n := 0
	var err error
	for ; nsum < len(bytes); nsum += n {
		n, err = file.Write(bytes[nsum:])
		if err != nil { // should also never happen
			return errors.New(CurrentFunctionName() + ":" + err.Error())
		}
	}
	return nil
}

// eof
