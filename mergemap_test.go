package mergemap

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestMerge(t *testing.T) {
	for _, tuple := range []struct {
		src      string
		dst      string
		expected string
	}{
		{
			src:      `{}`,
			dst:      `{}`,
			expected: `{}`,
		},
		{
			src:      `{"b":2}`,
			dst:      `{"a":1}`,
			expected: `{"a":1,"b":2}`,
		},
		{
			src:      `{"a":0}`,
			dst:      `{"a":1}`,
			expected: `{"a":0}`,
		},
		{
			src:      `{"a":{       "y":2}}`,
			dst:      `{"a":{"x":1       }}`,
			expected: `{"a":{"x":1, "y":2}}`,
		},
		{
			src:      `{"a":{"x":2}}`,
			dst:      `{"a":{"x":1}}`,
			expected: `{"a":{"x":2}}`,
		},
		{
			src:      `{"a":{       "y":7, "z":8}}`,
			dst:      `{"a":{"x":1, "y":2       }}`,
			expected: `{"a":{"x":1, "y":7, "z":8}}`,
		},
		{
			src:      `{"1": { "b":1, "2": { "3": {         "b":3, "n":[1,2]} }        }}`,
			dst:      `{"1": {        "2": { "3": {"a":"A",        "n":"xxx"} }, "a":3 }}`,
			expected: `{"1": { "b":1, "2": { "3": {"a":"A", "b":3, "n":[1,2]} }, "a":3 }}`,
		},
	} {
		var dst map[string]interface{}
		if err := json.Unmarshal([]byte(tuple.dst), &dst); err != nil {
			t.Error(err)
			continue
		}

		var src map[string]interface{}
		if err := json.Unmarshal([]byte(tuple.src), &src); err != nil {
			t.Error(err)
			continue
		}

		var expected map[string]interface{}
		if err := json.Unmarshal([]byte(tuple.expected), &expected); err != nil {
			t.Error(err)
			continue
		}

		got := Merge(dst, src)
		assert(t, expected, got)
	}
}

func TestMergeWithConfig(t *testing.T) {
	testCases := []struct {
		src      string
		dst      string
		config   map[string]interface{}
		expected string
	}{
		{
			src: `{"my-key": "new value"}`,
			dst: `{"my-key": "original value"}`,
			config: map[string]interface{}{
				"my-key": StrategyLastValue,
			},
			expected: `{"my-key": "new value"}`,
		},
		{
			src: `{"my-key": "new value"}`,
			dst: `{"my-key": "original value"}`,
			config: map[string]interface{}{
				"my-key": StrategyFirstValue,
			},
			expected: `{"my-key": "original value"}`,
		},
		{
			src: `{"my-key": {"my-sub": "new value"}}`,
			dst: `{"my-key": {"my-sub": "original value"}}`,
			config: map[string]interface{}{
				"my-key": map[string]interface{}{
					"my-sub": StrategyLastValue,
				},
			},
			expected: `{"my-key": {"my-sub": "new value"}}`,
		},
		{
			src: `{"my-key": {"my-sub": "new value"}}`,
			dst: `{"my-key": {"my-sub": "original value"}}`,
			config: map[string]interface{}{
				"my-key": map[string]interface{}{
					"my-sub": StrategyFirstValue,
				},
			},
			expected: `{"my-key": {"my-sub": "original value"}}`,
		},
	}

	for _, testCase := range testCases {
		var dst map[string]interface{}
		if err := json.Unmarshal([]byte(testCase.dst), &dst); err != nil {
			t.Error(err)
			continue
		}

		var src map[string]interface{}
		if err := json.Unmarshal([]byte(testCase.src), &src); err != nil {
			t.Error(err)
			continue
		}

		var expected map[string]interface{}
		if err := json.Unmarshal([]byte(testCase.expected), &expected); err != nil {
			t.Error(err)
			continue
		}

		got := MergeWithConfig(dst, src, testCase.config)
		assert(t, expected, got)
	}

}

func assert(t *testing.T, expected, got map[string]interface{}) {
	expectedBuf, err := json.Marshal(expected)
	if err != nil {
		t.Error(err)
		return
	}
	gotBuf, err := json.Marshal(got)
	if err != nil {
		t.Error(err)
		return
	}
	if bytes.Compare(expectedBuf, gotBuf) != 0 {
		t.Errorf("expected %s, got %s", string(expectedBuf), string(gotBuf))
		return
	}
}
