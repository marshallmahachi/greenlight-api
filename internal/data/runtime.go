package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {

	jsonValue := fmt.Sprintf("%d minutes", r)

	return []byte(strconv.Quote(jsonValue)), nil
}
