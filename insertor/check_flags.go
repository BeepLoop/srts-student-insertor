package insertor

import (
	"errors"
	"fmt"
)

func CheckFlags(limit *bool, maxLimit *int, args ...*string) error {

	if *limit && *maxLimit <= 0 {
		return errors.New("maxLimit must be greater than 0")
	}

	for _, arg := range args {
		if *arg == "" {
			return errors.New(fmt.Sprintf("Missing required flag: %s", *arg))
		}
	}

	return nil
}
