package mergemap

type mergeStrategy string

const (
	StrategyLastValue  mergeStrategy = "last_value"  // Use the last value found
	StrategyFirstValue mergeStrategy = "first_value" // Use the first value found

	StrategyMaxValue mergeStrategy = "max_value"
	StrategyMinValue mergeStrategy = "min_value"
)

func shouldUpdateValue(dst map[string]interface{}, key string, val interface{}, config map[string]interface{}) bool {
	strat := StrategyLastValue

	if tmp, ok := config[key].(mergeStrategy); ok {
		strat = tmp
	}

	switch strat {
	case StrategyLastValue:
		return true // this is the default
	case StrategyFirstValue:
		if _, alreadySet := dst[key]; alreadySet {
			return false
		}
	case StrategyMaxValue:
		if curVal, exists := dst[key]; exists {
			var current float64
			var newer float64

			current, ok := curVal.(float64)
			if !ok {
				return false
			}

			// incoming might be int, in that case we need to convert to float64
			newer, ok = val.(float64)
			if !ok {
				tmp, ok := val.(int)
				if !ok {
					// type is neither float64 or int
					return false
				}

				newer = float64(tmp)
			}

			return newer > current
		}

		return true
	case StrategyMinValue:
		if curVal, exists := dst[key]; exists {
			var current float64
			var newer float64

			current, ok := curVal.(float64)
			if !ok {
				return false
			}

			// incoming might be int, in that case we need to convert to float64
			newer, ok = val.(float64)
			if !ok {
				tmp, ok := val.(int)
				if !ok {
					// type is neither float64 or int
					return false
				}

				newer = float64(tmp)
			}

			return newer < current
		}
	}

	return true
}
