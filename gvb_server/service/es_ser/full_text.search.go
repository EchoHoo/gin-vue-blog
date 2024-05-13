package es_ser

import (
	"context"
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/olivere/elastic/v7"
	"github.com/russross/blackfriday"
	"github.com/sirupsen/logrus"
)

type SearchData struct {
	Key   string `json:"key"`   //文章ID
	Body  string `json:"body"`  //正文
	Slug  string `json:"slug"`  //包含文章ID
	Title string `json:"title"` //标题
}

func AsyncArticleByFullText(id, title, content string) error {
	indexList := GetSearchIndexDataByContent(id, title, content)
	bulk := global.ESClient.Bulk()
	for _, id := range indexList {
		req := elastic.NewBulkIndexRequest().Index(models.FullTextModel{}.Index()).Doc(id)
		bulk.Add(req)
	}
	result, err := bulk.Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return err
	}
	logrus.Infof("%s 添加成功，共%d 条数据", title, len(result.Succeeded()))
	return nil
}
func DeleteFullTextByArticleID(id string) {
	boolSearch := elastic.NewTermQuery("key", id)
	res, _ := global.ESClient.
		DeleteByQuery().
		Index(models.FullTextModel{}.Index()).
		Query(boolSearch).
		Do(context.Background())
	logrus.Infof("删除成功，共%d 条数据", res.Deleted)
}

// var data = "#hello world1\n\n#hello world2\n\n#hello world3\n\n"
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
			Key:   id,
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
