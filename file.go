package go_libs

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"
)

func FilenameWithoutSuffix(filename string) string {
	extension := filepath.Ext(filename)
	return filename[0 : len(filename)-len(extension)]
}

// ByteArray2File writes a byte array into a file. If required it does so in multiple steps.
// If all succeeds then nil is returned, otherwise an error.
func ByteArray2File(file *os.File, bytes []byte) error {
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

// ByteArray2ReponseWriter writes a byte array into a file. If required it does so in multiple steps.
// If all succeeds then nil is returned, otherwise an error.
func ByteArray2ReponseWriter(file http.ResponseWriter, bytes []byte) error {
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
