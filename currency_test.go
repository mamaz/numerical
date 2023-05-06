package numerical

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    uint
	Expected string
}

func TestConvertWithComma(t *testing.T) {
	data := []TestCase{
		{
			Input:    1,
			Expected: "1",
		},
		{
			Input:    12,
			Expected: "12",
		},
		{
			Input:    120,
			Expected: "120",
		},
		{
			Input:    1234,
			Expected: "1,234",
		},
		{
			Input:    12345,
			Expected: "12,345",
		},
		{
			Input:    123456,
			Expected: "123,456",
		},
		{
			Input:    1234567,
			Expected: "1,234,567",
		},
		{
			Input:    12345678,
			Expected: "12,345,678",
		},
		{
			Input:    123456789,
			Expected: "123,456,789",
		},
	}

	for _, testcase := range data {
		result, err := ConvertToCurrency(testcase.Input, ",")
		assert.Equal(t, nil, err)
		assert.Equal(t, testcase.Expected, result)
	}
}

func TestConvertWithUnderscore(t *testing.T) {
	result, err := ConvertToCurrency(1234, "_")
	assert.Nil(t, err)
	assert.Equal(t, "1_234", result)
}
