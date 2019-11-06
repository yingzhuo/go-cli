package cli

import (
	"os"
	"strconv"
)

type fileModeValue struct {
	val *os.FileMode
}

func (v *fileModeValue) Set(value string) error {
	val, err := strconv.ParseUint(value, 0, 32)
	if err != nil {
		return err
	}

	*v.val = os.FileMode(val)
	return nil
}

func (v *fileModeValue) String() string {
	return v.val.String()
}
