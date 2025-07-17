package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidRuntimeFormat = errors.New("invalid runtime value")

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {

	jsonValue := fmt.Sprintf("%d minutes", r)

	return []byte(strconv.Quote(jsonValue)), nil
}

func (r *Runtime) UnmarshalJSON(data []byte) error {

	unquotedJSONValue, err := strconv.Unquote(string(data))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unquotedJSONValue, " ")

	if len(parts) != 2 || parts[1] != "minutes" {
		return ErrInvalidRuntimeFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*r = Runtime(i)

	return nil
}
