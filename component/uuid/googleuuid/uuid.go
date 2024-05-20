package googleuuid

import (
	"github.com/google/uuid"
	uuid2 "h5travelotobackend/component/uuid"
)

type googleUUID struct {
}

func NewGoogleUUID() *googleUUID {
	return &googleUUID{}
}

func (g *googleUUID) Generate() (string, error) {
	newUUID, err := uuid.NewV7()
	if err != nil {
		return "", uuid2.ErrCannotGenerateUUID
	}
	return newUUID.String(), nil
}
