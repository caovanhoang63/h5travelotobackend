package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorized(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewCustomError(root error, msg string, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}
	return NewErrorResponse(
		errors.New(msg),
		msg,
		msg,
		key)
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}
	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func ErrDb(err error) *AppError {
	return NewFullErrorResponse(
		http.StatusInternalServerError,
		err,
		"Something went wrong with DB",
		err.Error(),
		"DB_ERROR")
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "Invalid request", err.Error(), "INVALID_REQUEST")
}

func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "something went wrong in the server", err.Error(), "INTERNAL_ERROR")
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot list %s", strings.ToLower(entity)),
		fmt.Sprintf("CANNOT_LIST_%s", strings.ToUpper(entity)))
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("CANNOT_DELETE_%s", strings.ToUpper(entity)))
}

func ErrEntityDeleted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Entity %s has been deleted", strings.ToLower(entity)),
		fmt.Sprintf("ENTITY_DELETED_%s", strings.ToUpper(entity)))
}

func ErrEntityNotFound(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Entity %s not found", strings.ToLower(entity)),
		fmt.Sprintf("ENTITY_NOT_FOUND_%s", strings.ToUpper(entity)))
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot create %s", strings.ToLower(entity)),
		fmt.Sprintf("CANNOT_CREATE_%s", strings.ToUpper(entity)))
}

func ErrNoPermission(err error) *AppError {
	return NewCustomError(
		err,
		"You have no permission",
		"NO_PERMISSION")
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot update %s", strings.ToLower(entity)),
		fmt.Sprintf("CANNOT_UPDATE_%s", strings.ToUpper(entity)))
}

func ErrCannotPublishMessage(topic string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot publish message to %s", topic),
		fmt.Sprintf("CANNOT_PUBLISH_%s", strings.ToUpper(topic)))
}

func ErrTooManyRequest(ip, api string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Too many request from %s to %s", ip, api),
		fmt.Sprintf("TOO_MANY_REQUEST"))
}

func ErrToCacheEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot cache %s", strings.ToLower(entity)),
		fmt.Sprintf("CANNOT_CACHE_%s", strings.ToUpper(entity)))
}

var DocumentNotFound = errors.New("document not found")
var RecordNotFound = errors.New("record not found")
var RateLimited = errors.New("rate limited")
