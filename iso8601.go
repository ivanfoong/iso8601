// The iso8601 package encodes and decodes time.Time in JSON in
// ISO 8601 format, without subsecond resolution or time zone info.
package iso8601

import "time"

const Format = "2006-01-02T15:04:05+08:00"
const jsonFormat = `"` + Format + `"`

type Time time.Time

func (it Time) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(it).Format(jsonFormat)), nil
}

func (it *Time) UnmarshalJSON(data []byte) error {
	t, e := time.Parse(`"` + time.RFC3339 + `"`, string(data))

	if e == nil {
		*it = Time(t)
	}

	return e;
}

func (it Time) String() string {
	return time.Time(it).String()
}
