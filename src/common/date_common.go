package common

import "time"

const (
	yyyyMmDd    = "2006-01-02"
	yyyymmddHhmmss = "2006-01-02 15:04:05"
)

type DateCommon struct {

}

func (d *DateCommon) ConvertFromDateStr(dateStr string) time.Time {
	date, _ := time.Parse(yyyyMmDd, dateStr)
	return date
}

func (d *DateCommon) ConvertFromDateTimeStr(datetimeStr string) time.Time {
	date, _ := time.Parse(yyyymmddHhmmss, datetimeStr)
	return date
}

func (d *DateCommon) ConvertFromDate(date time.Time) string {
	return date.Format(yyyyMmDd)
}

func (d *DateCommon) ConvertFromDateTime(date time.Time) string {
	return date.Format(yyyymmddHhmmss)
}