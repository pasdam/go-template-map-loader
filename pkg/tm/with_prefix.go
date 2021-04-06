package tm

import "strings"

func WithPrefix(prefix string, data map[string]interface{}) map[string]interface{} {
	if prefix == "" {
		return data
	}

	parts := strings.Split(prefix, ".")
	result := make(map[string]interface{}, 1)
	currentMap := result
	for i := 0; i < len(parts)-1; i++ {
		mm := make(map[string]interface{}, 1)
		currentMap[parts[i]] = mm
		currentMap = mm
	}
	currentMap[parts[len(parts)-1]] = data
	return result
}
