package soscrud

import (
	"encoding/json"
	"github.com/Alvarios/s-os/utils"
)

func Get(path string, ptr interface{}) error {
	file, err := sosutils.OpenFromProjectRoot(path)

	if err != nil {
		return err
	}

	err = json.Unmarshal(file, ptr)
	return err
}
