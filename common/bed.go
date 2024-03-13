package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Bed struct {
	Single int `json:"single" gorm:"column:single;"`
	Double int `json:"double" gorm:"column:double;"`
}

func (Bed) TableName() string { return "bed" }

// Scan scan value into Jsonb,
// decode jsonb in db into struct
// implements sql.Scanner interface
func (b *Bed) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("Failed to unmarshal JSONB value: %v", value))
	}
	var bed Bed
	if err := json.Unmarshal(bytes, &bed); err != nil {
		return err
	}
	*b = bed
	return nil
}

// Value return json value;
// encode struct to []byte aka jsonb
// ;implement driver.Valuer interface
func (b *Bed) Value() (driver.Value, error) {
	if b == nil {
		return nil, nil
	}
	return json.Marshal(b)
}
