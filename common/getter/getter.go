package getter

import "reflect"

// IsNil ...
func IsNil(source interface{}) bool {
	return source == nil || reflect.ValueOf(source).Kind() == reflect.Ptr && reflect.ValueOf(source).IsNil()
}

// GetValueAsString ...
func GetValueAsString(source map[string]interface{}, key string, defaultVal string) string {
	if IsNil(source[key]) {
		return defaultVal
	}
	return reflect.ValueOf(source[key]).Interface().(string)
}
