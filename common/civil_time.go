package common

import (
	"database/sql/driver"
	"errors"
	"strings"
	"time"
)

type CivilTime time.Time

func (c *CivilTime) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse("2006-01-02", value) //parse time
	if err != nil {
		return err
	}
	*c = CivilTime(t) //set result using the pointer
	return nil
}

func (c CivilTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(c).Format("2006-01-02") + `"`), nil
}

func (c *CivilTime) Value() (driver.Value, error) {
	if c == nil {
		return nil, nil
	}
	return time.Time(*c), nil
}

func (c *CivilTime) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	if v, ok := value.(time.Time); ok {
		*c = CivilTime(v)
		return nil
	}
	return errors.New("invalid Scan Source")
}
