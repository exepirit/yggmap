package adminapi

func rawRequest(requestType string) map[string]interface{} {
	return map[string]interface{}{
		"request":   requestType,
		"keepalive": true,
	}
}

func requestWithKey(requestType, key string) map[string]interface{} {
	return map[string]interface{}{
		"request":   requestType,
		"keepalive": true,
		"arguments": map[string]interface{}{
			"key": key,
		},
	}
}
