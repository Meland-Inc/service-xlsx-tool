package excel

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

var withStrictExcelConfig bool = false

func WithStrictExcelConfig(enable bool) {
	withStrictExcelConfig = enable
}

func captureExcelConfigErr(err error) {
	if !withStrictExcelConfig {
		return
	}

	// FIXME: don't panic
	panic(err)
}

var (
	errInvalidConfigSheet = errors.New("错误的配置表结构")
)

func parseIntValue(raw string) int {
	i, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		captureExcelConfigErr(err)
		return 0
	}
	return int(i)
}

func parseStringValue(raw string) string {
	return raw
}

func parseStringSliceValue(raw string) (rv []string) {
	if 0 == len(raw) {
		return
	}
	return strings.Split(raw, ",")
}

func parseStringSliceSliceValue(raw string) (rv [][]string) {
	if 0 == len(raw) {
		return
	}

	for _, part := range strings.Split(raw, ";") {
		rv = append(rv, parseStringSliceValue(part))
	}
	return rv
}

func parseFloatValue(raw string) float64 {
	f, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		captureExcelConfigErr(err)
		return 0.0
	}
	return f
}

func parseFloatSliceValue(raw string) (rv []float64) {
	if 0 == len(raw) {
		return
	}
	for _, part := range strings.Split(raw, ",") {
		rv = append(rv, parseFloatValue(part))
	}
	return rv
}

func parseFloatSliceSliceValue(raw string) (rv [][]float64) {
	if 0 == len(raw) {
		return
	}
	for _, part := range strings.Split(raw, ";") {
		rv = append(rv, parseFloatSliceValue(part))
	}
	return rv
}

func parseBoolValue(raw string) bool {
	return raw == "1"
}

func parseValueWithType(raw, fieldname, fieldtype string) (interface{}, bool) {
	if fieldname == "null" || fieldname == "" {
		// 不需要处理的字段
		return nil, false
	}

	switch fieldtype {
	case "int":
		return parseIntValue(raw), true
	case "int[]":
		return ParseIntSliceValue(raw), true
	case "int[][]":
		return ParseIntSliceSliceValue(raw), true
	case "string":
		return parseStringValue(raw), true
	case "string[]":
		return parseStringSliceValue(raw), true
	case "string[][]":
		return parseStringSliceSliceValue(raw), true
	case "float":
		return parseFloatValue(raw), true
	case "float[]":
		return parseFloatSliceValue(raw), true
	case "float[][]":
		return parseFloatSliceSliceValue(raw), true
	case "bool":
		return parseBoolValue(raw), true
	default:
		captureExcelConfigErr(fmt.Errorf("unknown fieldtype: %s", fieldtype))
		return raw, true
	}
}

// 解析一个配置 excel 表
//
// | 中文字段名称 |
// | 英文字段名称 |
// | 类型 hinting |
// | 字段值       |
func ParseExcelSheet(xlsx *excelize.File, sheet string) (rv []map[string]interface{}, err error) {
	rows := xlsx.GetRows(sheet)
	if len(rows) < 4 {
		return nil, errInvalidConfigSheet
	}

	fields, typeHintings := rows[1], rows[2]
	for idx, _ := range fields {
		fields[idx] = strings.TrimSpace(fields[idx])
		typeHintings[idx] = strings.TrimSpace(typeHintings[idx])
	}
	for i := 3; i < len(rows); i++ {
		row := rows[i]
		rowValue := map[string]interface{}{}

		//跳过空行
		if 0 == strings.Compare(row[0], "") {
			continue
		}

		for idx, field := range fields {
			value, keep := parseValueWithType(
				row[idx],
				field,
				typeHintings[idx],
			)
			if keep {
				rowValue[field] = value
			}
		}

		rv = append(rv, rowValue)
	}

	return rv, nil
}

func FirstSheet(xlsx *excelize.File) string {
	firstSheetId := int64(1e10)
	firstSheetName := ""
	for id, name := range xlsx.GetSheetMap() {
		if int64(id) < firstSheetId {
			firstSheetId = int64(id)
			firstSheetName = name
		}
	}
	if firstSheetName == "" {
		captureExcelConfigErr(fmt.Errorf("empty sheet"))
		return ""
	}
	return firstSheetName
}
