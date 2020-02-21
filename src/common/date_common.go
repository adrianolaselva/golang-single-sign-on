package common

import "time"

const (
	YYYY_MM_DD    		= "2006-01-02"
	YYYY_MM_DD_HH_MM_SS = "2006-01-02 15:04:05"
)

type DateCommon interface {
	ConvertFromDateStr(dateStr string) time.Time
	ConvertFromDateTimeStr(datetimeStr string) time.Time
	ConvertFromDate(date time.Time) string
	ConvertFromDateTime(date time.Time) string
}

type dateCommonImpl struct {

}

func NewDateCommon() *dateCommonImpl  {
	return &dateCommonImpl{}
}

func (d *dateCommonImpl) ConvertFromDateStr(dateStr string) time.Time {
	date, _ := time.Parse(YYYY_MM_DD, dateStr)
	return date
}

func (d *dateCommonImpl) ConvertFromDateTimeStr(datetimeStr string) time.Time {
	date, _ := time.Parse(YYYY_MM_DD_HH_MM_SS, datetimeStr)
	return date
}

func (d *dateCommonImpl) ConvertFromDate(date time.Time) string {
	return date.Format(YYYY_MM_DD)
}

func (d *dateCommonImpl) ConvertFromDateTime(date time.Time) string {
	return date.Format(YYYY_MM_DD_HH_MM_SS)
}