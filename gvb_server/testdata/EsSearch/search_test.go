package essearch

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
)

type SearchData struct {
	Body  string `json:"body"`  //正文
	Slug  string `json:"slug"`  //包含文章ID
	Title string `json:"title"` //标题
}

func GetSearchIndexDataByContent(id, title, content string) (SearchDataList []SearchData) {

	dataList := strings.Split(content, "\n")
	var isCode bool = false
	var body string
	var headList, bodyList []string
	headList = append(headList, getHeader(title))
	for _, v := range dataList {
		// #{1,6}
		// 判断一下是否是代码块
		if strings.HasPrefix(v, "```") {
			isCode = !isCode
		}
		if strings.HasPrefix(v, "#") && !isCode {
			headList = append(headList, v)
			if strings.TrimSpace(v) != "" {
				bodyList = append(bodyList, body)
			}
			body = ""
			continue
		}
		body += v
	}
	bodyList = append(bodyList, getBody(body))
	ln := len(headList)
	for i := 0; i < ln; i++ {
		SearchDataList = append(SearchDataList, SearchData{
			Title: headList[i],
			Body:  getBody(bodyList[i]),
			Slug:  id + getSlug(headList[i]),
		})
	}
	return SearchDataList
}
func getHeader(head string) string {
	head = strings.ReplaceAll(head, "#", "")
	head = strings.ReplaceAll(head, " ", "")
	return head
}
func getBody(body string) string {
	unsafe := blackfriday.MarkdownCommon([]byte(body))
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	return doc.Text()

}

func getSlug(slug string) string {
	return "#" + slug
}
func TestSearch(t *testing.T) {
	var data = "#hello world1\n\n#hello world2\n\n#hello world3\n\n"
	GetSearchIndexDataByContent("article/12", "test", data)
}
