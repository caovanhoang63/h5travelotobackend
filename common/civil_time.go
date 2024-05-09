package common

import (
	"database/sql/driver"
	"errors"
	"strings"
	"time"
)

type CivilDate time.Time

func (c *CivilDate) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse("01-02-2006", value) //parse time
	if err != nil {
		return err
	}
	*c = CivilDate(t) //set result using the pointer
	return nil
}

func (c CivilDate) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(c).Format("01-02-2006") + `"`), nil
}

func (c *CivilDate) Value() (driver.Value, error) {
	if c == nil {
		return nil, nil
	}
	return time.Time(*c), nil
}

func (c *CivilDate) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	if v, ok := value.(time.Time); ok {
		*c = CivilDate(v)
		return nil
	}
	return errors.New("invalid Scan Source")
}
