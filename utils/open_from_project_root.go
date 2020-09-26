package sosutils

import (
	"io/ioutil"
	"strings"
)

func OpenFromProjectRoot(path string) ([]byte, error) {
	slicePath := strings.Split(path, "/")

	if slicePath[0] != "" {
		rootPath, err := FindProjectRoot(slicePath[0])

		if err != nil {
			return nil, err
		}

		slicePath[0] = rootPath
	}

	return ioutil.ReadFile(strings.Join(slicePath, "/"))
}
