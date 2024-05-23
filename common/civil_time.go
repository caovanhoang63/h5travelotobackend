package common

import (
	"database/sql/driver"
	"errors"
	"math"
	"strings"
	"time"
)

const CivilDateFormat = "02-01-2006"

type CivilDate time.Time

func (c *CivilDate) String() string {
	return time.Time(*c).Format(CivilDateFormat)
}

func (c *CivilDate) After(d CivilDate) bool {
	return time.Time(*c).After(time.Time(d))
}

func (c *CivilDate) Before(d CivilDate) bool {
	return time.Time(*c).Before(time.Time(d))
}

func (c *CivilDate) DateDiff(d CivilDate) int {
	return int(math.Abs(time.Time(*c).Sub(time.Time(d)).Hours()) / 24)
}

func (c *CivilDate) ToString() string {
	return time.Time(*c).Format(CivilDateFormat)
}

func (c *CivilDate) IsEqual(d CivilDate) bool {
	return c.DateDiff(d) == 0
}

func (c *CivilDate) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse(CivilDateFormat, value) //parse time
	if err != nil {
		return err
	}
	*c = CivilDate(t) //set result using the pointer
	return nil
}

func (c CivilDate) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(c).Format(CivilDateFormat) + `"`), nil
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
