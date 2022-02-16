package mergemap

type mergeStrategy string

const (
	StrategyLastValue  mergeStrategy = "last_value"  // Use the last value found
	StrategyFirstValue mergeStrategy = "first_value" // Use the first value found
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
	}

	return true
}
