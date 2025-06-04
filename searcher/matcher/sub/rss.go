package sub

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	d "github.com/yshujie/goinaction/searcher/data"
	m "github.com/yshujie/goinaction/searcher/matcher"
)

// rssMatcher rss 格式数据匹配器
type rssMatcher struct{}

// init 函数，注册 rss matcher
func init() {
	var matcher rssMatcher
	m.Register("rss", matcher)
}

// Search 函数，根据关键词进行搜索
func (matcher rssMatcher) Search(feed *d.Feed, seatchTerm string) ([]*m.Result, error) {
	var results []*m.Result

	// 获取 feed 的文档数据
	doc, err := matcher.retrieve(feed)
	if err != nil {
		return nil, err
	}

	for _, channelItem := range doc.Channel.Item {
		// 根据 title 检索
		matched, err := regexp.MatchString(seatchTerm, channelItem.Title)
		if err != nil {
			return nil, err
		}
		if matched {
			results = append(results, &m.Result{
				Field:   "Title",
				Content: channelItem.Title,
			})
		}

		// 根据 Description 检索
		matched, err = regexp.MatchString(seatchTerm, channelItem.Description)
		if err != nil {
			return nil, err
		}
		if matched {
			results = append(results, &m.Result{
				Field:   "Description",
				Content: channelItem.Description,
			})
		}
	}

	return results, nil
}

// retrieve 检索数据
func (matcher rssMatcher) retrieve(feed *d.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New("No rss feed uri provided")
	}

	// Retrieve the rss feed document from the web.
	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}

	// Close the response once we return from the function.
	defer resp.Body.Close()

	// Check the status code for a 200 so we know we have received a
	// proper response.
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Response Error %d\n", resp.StatusCode)
	}

	// Decode the rss feed document into our struct type.
	// We don't need to check for errors, the caller can do this.
	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}

type (
	// item defines the fields associated with the item tag
	// in the rss document.
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}

	// image defines the fields associated with the image tag
	// in the rss document.
	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	// channel defines the fields associated with the channel tag
	// in the rss document.
	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}

	// rssDocument defines the fields associated with the rss document.
	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)
