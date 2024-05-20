package uuid

import "errors"

type Uuid interface {
	Generate() (string, error)
}

var (
	ErrCannotGenerateUUID = errors.New("cannot generate uuid")
)
