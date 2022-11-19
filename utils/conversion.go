package utils

import (
	"log"
	"math"
	"strconv"
	"strings"
)

// StrToInt 字符串转int
func StrToInt(number string, defaultNum ...int) int {
	num, err := strconv.Atoi(number)
	if err != nil {
		if len(defaultNum) > 0 {
			return defaultNum[0]
		}
		return 0
	}
	return num
}

// StrToInt64 StrNum需要转换的字符串
// defaultNum默认值
// String类型转int64
func StrToInt64(strNum string, defaultNum ...int64) int64 {
	num, err := strconv.ParseInt(strNum, 10, 64)
	if err != nil {
		if len(defaultNum) > 0 {
			return defaultNum[0]
		} else {
			return 0
		}
	}
	return num
}

// StrToInt32 StrNum需要转换的字符串
// defaultNum默认值
// String类型转int64
func StrToInt32(strNum string, defaultNum ...int32) int32 {
	num, err := strconv.ParseInt(strNum, 10, 32)
	if err != nil {
		if len(defaultNum) > 0 {
			return defaultNum[0]
		} else {
			return 0
		}
	}
	return int32(num)
}

// StrToFloat32 String转float32位
// strFloat 需要转的字符串
// defaultFloat 默认值，若没有出现转换异常直接返回0.0
func StrToFloat32(strFloat string, defaultFloat ...float32) float32 {
	float, err := strconv.ParseFloat(strFloat, 32)
	if err != nil {
		if len(defaultFloat) > 0 {
			return defaultFloat[0]
		} else {
			return 0.0
		}
	}
	return float32(float)
}

// StrToFloat64 String转float64位
// strFloat 需要转的字符串
// defaultFloat 默认值，若没有出现转换异常直接返回0.0
func StrToFloat64(strFloat string, defaultFloat ...float64) float64 {
	float, err := strconv.ParseFloat(strFloat, 36)
	if err != nil {
		if len(defaultFloat) > 0 {
			return defaultFloat[0]
		} else {
			return 0.0
		}
	}
	return float
}

// IntToStr int类型转字符串
// num int类型的数据，不区分int64或者int类型
func IntToStr(num interface{}) string {
	var str string
	switch num.(type) {
	case int:
		str = strconv.Itoa(num.(int))
	case int64:
		str = strconv.FormatInt(num.(int64), 10)
	case int32:
		str = strconv.FormatInt(int64(num.(int32)), 10)
	case string:
		str = num.(string)
	default:
		str = ""
	}
	return str
}

func FloatToInt(num float64) int {
	return int(num)
}

func StringAddress(str string) *string {
	return &str
}
func Int64Address(str int64) *int64 {
	return &str
}

// BinDec Binary to decimal
func BinDec(b string) (n int64) {
	s := strings.Split(b, "")
	l := len(s)
	i := 0
	d := float64(0)
	for i = 0; i < l; i++ {
		f, err := strconv.ParseFloat(s[i], 10)
		if err != nil {
			log.Println("Binary to decimal error:", err.Error())
			return -1
		}
		d += f * math.Pow(2, float64(l-i-1))
	}
	return int64(d)
}
