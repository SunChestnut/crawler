package service

// 使用内存去重
var visitedUrls = make(map[string]bool)

func IsDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
