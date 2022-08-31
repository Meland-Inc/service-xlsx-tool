package excel

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cast"
)

func StringToString(v interface{}) string {
	s, ok := v.(string)
	if !ok {
		return ""
	}
	return s
}

func BoolToBool(v interface{}) bool {
	b, ok := v.(bool)
	if !ok {
		return false
	}
	return b
}

func IntToInt8(v interface{}) int8 {
	i, ok := v.(int)
	if !ok {
		return 0
	}
	return int8(i)
}

func IntToInt16(v interface{}) int16 {
	i, ok := v.(int)
	if !ok {
		return 0
	}
	return int16(i)
}

func IntToInt32(v interface{}) int32 {
	i, ok := v.(int)
	if !ok {
		return 0
	}
	return int32(i)
}

func IntToUint32(v interface{}) uint32 {
	i, ok := v.(int)
	if !ok {
		return 0
	}
	return uint32(i)
}

func IntToInt64(v interface{}) int64 {
	i, ok := v.(int)
	if !ok {
		return 0
	}
	return int64(i)
}

func IntToString(v interface{}) string {
	if v == nil {
		return ""
	}
	switch v := v.(type) {
	case nil:
		return ""
	case int, int64:
		return fmt.Sprintf("%d", v)
	default:
		panic(fmt.Sprintf("unsupported type: %v", v))
	}
}

func IntSliceToString(v interface{}) (str string) {
	is, ok := v.([]int)
	if !ok {
		return str
	}

	for _, i := range is {
		str += fmt.Sprintf("%d,", i)
	}
	if str != "" {
		str = str[:len(str)-1]
	}

	return str
}

func IntSliceToJsonStr(v interface{}) (str string) {
	is, ok := v.([]int)
	if !ok {
		panic(fmt.Sprintf("unsupported type: %v", v))
	}

	var arr []int

	for _, i := range is {
		arr = append(arr, int(i))
	}

	if len(arr) == 1 && arr[0] == 0 {
		return "[]"
	}
	rv, err := json.Marshal(arr)
	if err != nil {
		panic(err)
	}

	return string(rv)
}

func IntSliceToBinaryCode(v interface{}) (rv int32) {
	is, ok := v.([]int)
	if !ok {
		return 0
	}

	for _, i := range is {
		rv |= (1 << uint(i))
	}

	return rv
}

func IntSliceToInt32Slice(v interface{}) (rv []int32) {
	is, ok := v.([]int)
	if !ok {
		return nil
	}

	rv = []int32{}
	for _, i := range is {
		rv = append(rv, int32(i))
	}

	return
}

func IntSliceToInt64Slice(v interface{}) (rv []int64) {
	is, ok := v.([]int)
	if !ok {
		return nil
	}

	rv = []int64{}
	for _, i := range is {
		rv = append(rv, int64(i))
	}

	return
}

func IntSliceToStringSlice(v interface{}) (rv []string) {
	is, ok := v.([]int)
	if !ok {
		return nil
	}

	rv = []string{}
	for _, i := range is {
		rv = append(rv, fmt.Sprintf("%d", i))
	}

	return
}

func StringSliceToStringSlice(v interface{}) (rv []string) {
	rv, ok := v.([]string)
	if !ok {
		return nil
	}

	return
}

func StringSliceToString(v interface{}) string {
	rv, ok := v.([]string)
	if !ok {
		return ""
	}

	var str string
	for _, s := range rv {
		if str == "" {
			str = s
			continue
		}
		str += fmt.Sprintf(",%s", s)
	}

	return str
}

func StringSliceSliceToString(v interface{}) string {
	sss, ok := v.([][]string)
	if !ok {
		return ""
	}

	var rvStr string
	for index, ss := range sss {
		var str string
		for idx, s := range ss {
			if idx > 0 {
				str += ","
			}
			str += s
		}
		if index > 0 {
			rvStr += ";"
		}
		rvStr += str
	}
	return rvStr
}

func StringToInt32Slice(v interface{}, sep string) (slice []int32) {
	str, ok := v.(string)
	if !ok {
		return
	}

	splitArray := strings.Split(str, sep)
	for _, s := range splitArray {
		num := cast.ToInt32(s)
		slice = append(slice, num)
	}
	return
}
func StringToStringSlice(v interface{}, sep string) (slice []string) {
	str, ok := v.(string)
	if !ok {
		return
	}

	return strings.Split(str, sep)
}
