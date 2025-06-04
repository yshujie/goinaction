package sub

import (
	"log"

	d "github.com/yshujie/goinaction/searcher/data"
	m "github.com/yshujie/goinaction/searcher/matcher"
)

// defaultMatcher 默认匹配器
type defaultMatcher struct{}

// init 函数，注册 rss matcher
func init() {
	log.Println("in default matcher init")

	var matcher defaultMatcher
	m.Register("default", matcher)
}

// Search 函数，根据关键词进行搜索
func (matcher defaultMatcher) Search(feed *d.Feed, seatchTerm string) ([]*m.Result, error) {
	return nil, nil
}
