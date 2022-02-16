package mergemap

import (
	"reflect"
)

var (
	MaxDepth = 32
)

func MergeWithConfig(dst, src, config map[string]interface{}) map[string]interface{} {
	return merge(dst, src, 0, config)
}

// Merge recursively merges the src and dst maps. Key conflicts are resolved by
// preferring src, or recursively descending, if both src and dst are maps.
func Merge(dst, src map[string]interface{}) map[string]interface{} {
	return merge(dst, src, 0, nil)
}

func merge(dst, src map[string]interface{}, depth int, config map[string]interface{}) map[string]interface{} {
	if depth > MaxDepth {
		panic("too deep!")
	}
	for key, srcVal := range src {
		if dstVal, ok := dst[key]; ok {
			srcMap, srcMapOk := mapify(srcVal)
			dstMap, dstMapOk := mapify(dstVal)
			if srcMapOk && dstMapOk {
				subConfig, ok := config[key].(map[string]interface{})
				if !ok {
					subConfig = nil
				}
				srcVal = merge(dstMap, srcMap, depth+1, subConfig)
			}
		}

		if config != nil {
			var doUpdate bool
			if _, exists := config[key]; exists {
				doUpdate = shouldUpdateValue(dst, key, srcVal, config)
			}

			if !doUpdate {
				continue
			}
		}

		dst[key] = srcVal

	}
	return dst
}

func mapify(i interface{}) (map[string]interface{}, bool) {
	value := reflect.ValueOf(i)
	if value.Kind() == reflect.Map {
		m := map[string]interface{}{}
		for _, k := range value.MapKeys() {
			m[k.String()] = value.MapIndex(k).Interface()
		}
		return m, true
	}
	return map[string]interface{}{}, false
}
