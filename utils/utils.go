package utils

// InList 判断key是否存在与列表中
func InList(key string, list []string) bool {
	for _, v := range list {
		if key == v {
			return true
		}
	}
	return false
}
