package utils

import (
	"fmt"
	"runtime"
	"strings"
)

// HandleError ...
func HandleError(errIn error) (errOut error) {

	// notice that we're using 1, so it will actually log where
	// the error happened, 0 = this function, we don't want that.
	_, fn, line, _ := runtime.Caller(1)

	paths := strings.Split(fn, "/")

	rootPath := strings.ToLower("go-sqs-worker")

	var file string
	for i := range paths {
		if paths[i] == rootPath {
			for j := (i + 1); j < len(paths); j++ {
				file = fmt.Sprintf("%s/%s", file, paths[j])
			}
			break
		}
	}

	errOut = fmt.Errorf("%s:%d - %v", file, line, errIn)

	return errOut
}
