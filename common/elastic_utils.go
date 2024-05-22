package common

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"strconv"
)

func EsFloatToGoFloat(float types.Float64) (float64, error) {
	valB, err := float.MarshalJSON()
	a, err := strconv.ParseFloat(string(valB), 10)
	if err != nil {
		return 0, err
	}
	return a, nil
}
