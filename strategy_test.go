package mergemap

import (
	"encoding/json"
	"testing"
)

func TestShouldUpdateValue(t *testing.T) {
	testCases := []struct {
		dst      string
		key      string
		val      string
		cfg      FieldConfig
		expected bool
	}{
		{
			dst:      `{"key1": "value1"}`,
			key:      "key1",
			val:      "value2",
			cfg:      FieldConfig{Strategy: StrategyLastValue},
			expected: true,
		},
		{
			dst:      `{"key1": "value1"}`,
			key:      "key1",
			val:      "value2",
			cfg:      FieldConfig{Strategy: StrategyFirstValue},
			expected: false,
		},
	}

	for _, testCase := range testCases {
		var dst map[string]interface{}
		if err := json.Unmarshal([]byte(testCase.dst), &dst); err != nil {
			t.Error(err)
		}

		returnVal := shouldUpdateValue(dst, testCase.key, testCase.val, testCase.cfg)

		if returnVal != testCase.expected {
			t.Errorf("expected: %t, got: %t", testCase.expected, returnVal)
		}

	}
}