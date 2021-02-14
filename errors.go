package orgojson

import "fmt"

var (
	ErrNotFound = fmt.Errorf("the key was not found")
	ErrWrongType = fmt.Errorf("the value was of the wrong type")
)
