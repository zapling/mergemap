package mergemap

import (
	"encoding/json"
	"testing"
)

func TestShouldUpdateValue(t *testing.T) {
	testCases := []struct {
		dst      string
		key      string
		val      interface{}
		cfg      map[string]interface{}
		expected bool
	}{
		{
			dst: `{"key1": "value1"}`,
			key: "key1",
			val: "value2",
			cfg: map[string]interface{}{
				"key1": StrategyLastValue,
			},
			expected: true,
		},
		{
			dst: `{"key1": "value1"}`,
			key: "key1",
			val: "value2",
			cfg: map[string]interface{}{
				"key1": StrategyFirstValue,
			},
			expected: false,
		},
		{
			dst: `{"key1": 1}`,
			key: "key1",
			val: 2,
			cfg: map[string]interface{}{
				"key1": StrategyMaxValue,
			},
			expected: true,
		},
		{
			dst: `{"key1": 2}`,
			key: "key1",
			val: 1,
			cfg: map[string]interface{}{
				"key1": StrategyMinValue,
			},
			expected: true,
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
