package util

import (
	"database/sql"
	"strings"
	"time"
)

func WildCardString(keyword string) string {
	if keyword != "" {
		return "%" + keyword + "%"
	} else {
		return "%"
	}
}

func StringToDate(value string) time.Time {
	var dateValue time.Time
	dateValue, _ = time.Parse(DateLayoutYMD, value)
	return dateValue
}

func NewNullableString(value string) sql.NullString {
	if len(value) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: value,
		Valid:  true,
	}
}

func NewNullableDate(value time.Time) sql.NullTime {
	if value.IsZero() {
		return sql.NullTime{}
	}
	return sql.NullTime{
		Time:  value,
		Valid: true,
	}
}

func StringToArray(value string) []string {
	var result []string
	if len(value) > 0 {
		result = strings.Split(value, `,`)
	}
	return result
}

func ArrayToString(value []string) string {
	var result string
	if len(value) > 0 {
		result = strings.Join(value, `,`)
	}
	return result
}
