package matcher

import (
	d "github.com/yshujie/goinaction/searcher/data"
)

// Result Search 后返回的结果
type Result struct {
	Field   string
	Content string
}

// Matcher 匹配器接口，定义匹配器行为
type Matcher interface {
	Search(feed *d.Feed, seatchTerm string) ([]*Result, error)
}
