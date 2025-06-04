package searcher

import (
	"log"
	"sync"

	d "github.com/yshujie/goinaction/searcher/data"
	m "github.com/yshujie/goinaction/searcher/matcher"
)

// run 函数，执行搜索功能
func run(seatchTerm string) {
	// 拉取数据源
	feeds, err := d.RetrieveFeeds()
	if err != nil {
		log.Fatal("retrieve feeds fail, error: ", err)
	}

	// 创建 waitGroup，设置根据 feeds 搜索进行等待
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(feeds))

	// 遍历 feeds，选取对应的 matcher 进行搜索，将搜索结果存入 results 中
	results := make(chan *m.Result)
	for _, feed := range feeds {
		// 选择数据匹配器
		matcher := m.SelectMatcher(feed.Type)

		// 启用 goroutine 去执行搜索
		go func(matcher m.Matcher, feed *d.Feed) {
			m.Match(matcher, feed, seatchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// 启动搜索等待协程，搜索结束后关闭 results 通道
	go func() {
		// 等待 waitGroup 结束
		waitGroup.Wait()

		close(results)
	}()

	// 展示查询结构
	display(results)
}

// display 展示查询结构
func display(results chan *m.Result) {
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
