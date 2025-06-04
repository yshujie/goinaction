package matcher

import (
	"log"

	d "github.com/yshujie/goinaction/searcher/data"
)

// matchers，匹配器池
var matchers = make(map[string]Matcher)

// Register 匹配器注册器
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher is already registered")
	}

	matchers[feedType] = matcher
	log.Println("Register", feedType, "matcher")
}

// SelectMatcher 根据源数据类型获取对应的 matcher
func SelectMatcher(dataType string) Matcher {
	matcher, exists := matchers[dataType]
	if !exists {
		matcher = matchers["default"]
	}

	return matcher
}

// Match 函数，匹配器门面，对外提根据关键词供搜索功能
// 接收一个匹配器实例、数据源和搜索词，返回结果到结果通道
func Match(matcher Matcher, feed *d.Feed, searchTerm string, results chan<- *Result) {
	// 使用 matcher 搜索数据
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println("search err:", err)
		return
	}

	// 将搜索结果写入 result 通道
	for _, result := range searchResults {
		results <- result
	}
}
