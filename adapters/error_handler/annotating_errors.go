package error_handler

import (
	"fmt"
	"github.com/pkg/errors"
)

func AnnotatingError(err error, message string) error {
	if err != nil {
		fmt.Println(errors.Wrap(err, message))
		return errors.Wrap(err, message)
	}
	return nil
}
