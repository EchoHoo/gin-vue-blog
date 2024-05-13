package mth

import (
	"fmt"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
)

func TestMarkdownToHTML(t *testing.T) {
	unsafe := blackfriday.MarkdownCommon([]byte("###你好\n```python\nprint('你好')\n``` \n\n<script>alert('hello')</script>\n\n"))

	fmt.Println(string(unsafe))
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	doc.Find("script").Remove()
	fmt.Println(doc.Text())

}
