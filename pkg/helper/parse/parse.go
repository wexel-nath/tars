package parse

import (
	"strconv"
	"time"
)

func StringToFloat(value string) (float64, error) {
	return strconv.ParseFloat(value, 64)
}

func MustGetFloat(value string) float64 {
	f, _ := StringToFloat(value)
	return f
}

func MustParseFloat(value interface{}) float64 {
	floatVal, err := strconv.ParseFloat(BytesAsString(value), 64)
	if err != nil {
		panic(err)
	}

	return floatVal
}

func FloatPointer(value interface{}) *float64 {
	if value == nil {
		return nil
	}

	v := MustParseFloat(value)
	return &v
}

func StringToInt(value string) (int64, error) {
	return strconv.ParseInt(value, 10, 64)
}

func MustGetInt(value string) int64 {
	i, _ := StringToInt(value)
	return i
}

func IntAsString(value interface{}) string {
	return strconv.FormatInt(value.(int64), 10)
}

func BytesAsString(value interface{}) string {
	return string(value.([]byte))
}

func StringPointer(value interface{}) *string {
	if value == nil {
		return nil
	}

	v := value.(string)
	return &v
}

func TimePointer(value interface{}) *time.Time {
	if value == nil {
		return nil
	}

	v := value.(time.Time)
	return &v
}
