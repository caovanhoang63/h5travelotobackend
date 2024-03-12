package uploadmodel

import (
	"errors"
	"h5travelotobackend/common"
)

const EntityName = "Upload"

type Upload struct {
	common.SqlModel `json:",inline"`
	common.Image    `json:",inline"`
}

func (Upload) TableName() string {
	return "uploads"
}

var (
	ErrFileTooLarge = common.NewCustomError(
		errors.New("file too large"),
		"file too large",
		"Error_File_Too_Large",
	)
)

func ErrFileIsNotImage(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"file is not an image",
		"Error_File_Is_Not_Image",
	)
}
func ErrCannotSaveFile(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"cannot save file",
		"Error_Cannot_Save_File",
	)

}
