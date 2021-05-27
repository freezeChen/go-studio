package jsontime

import "time"

const (
	TIMEFORMAT string = "2006-01-02 15:04:05"
)

type JsonTime time.Time

func (t *JsonTime) GobDecode(data []byte) error {
	if (string(data)) == `""` {
		return nil
	}
	now, err := time.ParseInLocation(`"`+TIMEFORMAT+`"`, string(data), time.Local)
	*t = JsonTime(now)
	return err
}

func (t *JsonTime) GobEncode() ([]byte, error) {
	b := make([]byte, 0, len(TIMEFORMAT)+2)
	//当时间为0值时返回空字符串
	if time.Time(*t).IsZero() {
		b = append(b, '"', '"')
		return b, nil
	}

	b = append(b, '"')
	b = time.Time(*t).AppendFormat(b, TIMEFORMAT)
	b = append(b, '"')
	return b, nil
}

func (t *JsonTime) UnmarshalJSON(data []byte) (err error) {
	dataStr := string(data)
	if (dataStr) == `""` {
		return nil
	}
	now, err := time.ParseInLocation(`"`+TIMEFORMAT+`"`, string(data), time.Local)
	*t = JsonTime(now)
	return
}

func (t JsonTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TIMEFORMAT)+2)
	//当时间为0值时返回空字符串
	if time.Time(t).IsZero() {
		b = append(b, '"', '"')
		return b, nil
	}
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TIMEFORMAT)
	b = append(b, '"')
	return b, nil
}

func (t JsonTime) String() string {
	if time.Time(t).IsZero() {
		return ""
	}
	return time.Time(t).Format(TIMEFORMAT)
}

func (t JsonTime) Time() time.Time {
	return time.Time(t)
}

func Now() JsonTime {
	return JsonTime(time.Now())
}

func From(timeStr string) JsonTime {
	t, _ := time.ParseInLocation(TIMEFORMAT, timeStr, time.Local)
	return JsonTime(t)
}
