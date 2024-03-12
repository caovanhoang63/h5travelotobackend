package uploadprovider

import (
	"context"
	"h5travelotobackend/common"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
