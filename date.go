package datatypes

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type Date struct {
	time.Time
}

func (date *Date) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*date = Date{nullTime.Time}
	return
}

func (date *Date) UnmarshalParam(param string) error {
	date.Scan(param)
	return nil
}

func (date Date) Value() (driver.Value, error) {
	y, m, d := date.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, date.Location()), nil
}

// GormDataType gorm common data type
func (date Date) GormDataType() string {
	return "date"
}

func (date Date) GobEncode() ([]byte, error) {
	return date.Time.GobEncode()
}

func (date *Date) GobDecode(b []byte) error {
	return date.Time.GobDecode(b)
}

func (date Date) MarshalJSON() ([]byte, error) {
	println(date.Time.Format(time.RFC3339))
	return []byte("\"" + date.Time.Format(time.RFC3339) + "\""), nil
}

func (date *Date) UnmarshalJSON(b []byte) error {
	var y, mo, d, h, m, s, n int
	dateString := string(b)

	_, _ = fmt.Sscanf(dateString, "\"%04d-%02d-%02dT%02d:%02d:%02d.%d", &y, &mo, &d, &h, &m, &s, &n)
	date.Time = time.Date(y, time.Month(mo), d, h, m, s, n, date.Location())

	print(date.Time)

	return nil
}
