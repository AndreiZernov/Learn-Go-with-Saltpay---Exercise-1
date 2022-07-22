package error_handler

import (
	"fmt"
	"github.com/pkg/errors"
)

func AnnotatingError(err error, message string) {
	if err != nil {
		fmt.Println(errors.Wrap(err, message))
	}
}
