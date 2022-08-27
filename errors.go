package cyclicsort

import "errors"

var (
	ErrValueNegative   error = errors.New("slice has a negative value")
	ErrValueBig        error = errors.New("slice has a value bigger or equal to the size of the slice")
	ErrValueDuplicated error = errors.New("slice has duplicated values")
)
