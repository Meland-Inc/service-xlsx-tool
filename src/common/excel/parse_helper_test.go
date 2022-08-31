package excel

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseValueWithType(t *testing.T) {
	WithStrictExcelConfig(true)

	for _, c := range []struct {
		raw           string
		fieldname     string
		fieldtype     string
		expectedValue interface{}
		expectedKeep  bool
	}{
		{"foobar", "null", "int", nil, false},
		{"foobar", "string", "string", "foobar", true},
		{"foo,bar", "string[]", "string[]", []string{"foo", "bar"}, true},
		{"42", "int", "int", 42, true},
		{"4,2", "int[]", "int[]", []int{4, 2}, true},
		{"4,2;42", "int[][]", "int[][]", [][]int{{4, 2}, {42}}, true},
		{"4.2", "float", "float", 4.2, true},
		{"4.2,42", "float[]", "float[]", []float64{4.2, 42}, true},
		{"4.2,42;0.42", "float[][]", "float[][]", [][]float64{{4.2, 42}, {0.42}}, true},
		{"1", "bool", "bool", true, true},
		{"0", "bool", "bool", false, true},
	} {
		valueParsed, valueShouldKeep := parseValueWithType(c.raw, c.fieldname, c.fieldtype)

		require.Equal(t, c.expectedValue, valueParsed)
		require.Equal(t, c.expectedKeep, valueShouldKeep)
	}
}
