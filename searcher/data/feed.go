package data

import (
	"encoding/json"
	"os"
)

// 数据文件地址
const dataFile = "data/data.json"

// Feed 需要处理的数据源结构
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds 读取并返回源数据
func RetrieveFeeds() ([]*Feed, error) {
	// 打开数据文件
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// 设定调用结束后关闭文件
	defer file.Close()

	// 使用 JSON 解析器读取源数据
	var feeds []*Feed
	if err := json.NewDecoder(file).Decode(&feeds); err != nil {
		return nil, err
	}

	// 返回读取结构
	return feeds, nil
}
