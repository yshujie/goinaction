package goroutine_test

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func MirroredQuery() string {
	urls := []string{
		"google.com",
		"github.com",
		"bing.com",
	}

	// 创建有缓冲的通道
	responses := make(chan string, len(urls))

	for _, url := range urls {
		go request(url, responses)
	}

	// 在有缓冲通道中读取第一个元素，就返回
	return <-responses
}

// request 函数，请求 url 并将结果写入 responses 通道
// responses 是一个只写型通道
func request(url string, responses chan<- string) {
	resp, err := http.Get(ensureScheme(url))
	if err != nil {
		log.Fatalln("请求出错，错误信息：", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("解析出错，错误信息：", err)
		return
	}

	responses <- string(body)
}

func ensureScheme(raw string) string {
	if !strings.HasPrefix(raw, "http://") && !strings.HasPrefix(raw, "https://") {
		raw = "https://" + raw
	}
	// 顺便规范化（可选）
	if _, err := url.Parse(raw); err != nil {
		return "https://example.com"
	}
	return raw
}
