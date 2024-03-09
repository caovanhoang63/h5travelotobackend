package tokenprovider

import (
	"errors"
	"h5travelotobackend/common"
	"time"
)

type Provider interface {
	Generate(data TokenPayload, expiry int) (*Token, error)
	Validate(token string) (*TokenPayload, error)
}

var (
	ErrNotFound = common.NewCustomError(
		errors.New("token not found"),
		"token not found",
		"ERROR_NOT_FOUND")

	ErrEncodingToken = common.NewCustomError(
		errors.New("error encoding the token"),
		"error encoding the token",
		"ERROR_ENCODING_TOKEN")

	ErrInvalidToken = common.NewCustomError(
		errors.New("invalid token provided"),
		"invalid token provided",
		"ERROR_INVALID_TOKEN",
	)
)

type Token struct {
	Token  string    `json:"Token"`
	Create time.Time `json:"created"`
	Expiry int       `json:"expiry"`
}

type TokenPayload struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
}
