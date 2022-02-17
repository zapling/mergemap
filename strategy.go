package mergemap

type MergeStrategy string

const (
	StrategyLastValue  MergeStrategy = "last_value"
	StrategyFirstValue MergeStrategy = "first_value"

	StrategyMaxValue MergeStrategy = "max_value"
	StrategyMinValue MergeStrategy = "min_value"
)

type mergeStrategyFunc func(dst map[string]interface{}, key string, value interface{}) bool

// DefaultMergeStrategies are the default supported merge strategies.
// If any other behaviour is desiered you could override this variable or extend it.
var DefaultMergeStrategies = map[MergeStrategy]mergeStrategyFunc{
	StrategyFirstValue: isTheFirstValue,
	StrategyLastValue:  isTheLastValue,
	StrategyMaxValue:   isTheMaximumValue,
	StrategyMinValue:   isTheMinimumValue,
}

// isTheLastValue is the default strategy, the latest value will always override if no other
// strategy is set
func isTheLastValue(dst map[string]interface{}, key string, value interface{}) bool {
	return true
}

// isTheFirstValue check whatever the key is already set in the destination map
func isTheFirstValue(dst map[string]interface{}, key string, val interface{}) bool {
	if _, alreadySet := dst[key]; alreadySet {
		return false
	}

	return true
}

// isTheMaximumValue checks whatever the value is higher than what is already set
func isTheMaximumValue(dst map[string]interface{}, key string, value interface{}) bool {
	if currentValue, alreadySet := dst[key]; alreadySet {
		var currentValueTyped float64
		var newValueTyped float64

		currentValueTyped, ok := currentValue.(float64)
		if !ok {
			return false
		}

		newValueTyped, ok = value.(float64)
		if !ok {
			// the value could be int, I think
			tmp, ok := value.(int)
			if !ok {
				// not it, we can't recover
				return false
			}

			newValueTyped = float64(tmp)
		}

		return newValueTyped > currentValueTyped
	}

	return true
}

// isTheMinimumValue checks whatever the value is lower then what is already set
func isTheMinimumValue(dst map[string]interface{}, key string, value interface{}) bool {
	if currentValue, alreadySet := dst[key]; alreadySet {
		var currentValueTyped float64
		var newValueTyped float64

		currentValueTyped, ok := currentValue.(float64)
		if !ok {
			return false
		}

		newValueTyped, ok = value.(float64)
		if !ok {
			// the value could be int, I think
			tmp, ok := value.(int)
			if !ok {
				// not it, we can't recover
				return false
			}

			newValueTyped = float64(tmp)
		}

		return newValueTyped < currentValueTyped
	}

	return true
}

// shouldUpdateValue check if there is any merge strategy set for the current key, and if there is
// it check the value against that strategy function to determine if the value should be
// updated or not.
func shouldUpdateValue(dst map[string]interface{}, key string, val interface{}, config map[string]interface{}) bool {
	strategy, ok := config[key].(MergeStrategy)
	if ok {
		if strategyFunc, exists := DefaultMergeStrategies[strategy]; exists {
			return strategyFunc(dst, key, val)
		}
	}

	return isTheLastValue(dst, key, val)
}
