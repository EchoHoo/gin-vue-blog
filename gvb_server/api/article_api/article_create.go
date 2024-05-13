package article_api

import (
	"gvb_server/gvb_server/global"
	"gvb_server/gvb_server/models"
	"gvb_server/gvb_server/models/ctype"
	"gvb_server/gvb_server/models/res"
	"gvb_server/gvb_server/service/es_ser"
	"gvb_server/gvb_server/utils/jwts"
	"math/rand"
	"strings"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
)

type ArticleRequest struct {
	Title   string `json:"title" binding:"required" msg:"文章标题不能为空"`   //文章标题
	Abstact string `json:"abstract"`                                  //文章简介
	Content string `json:"content" binding:"required" msg:"文章内容不能为空"` //文章内容

	Category string `json:"category"` //文章分类
	Source   string `json:"source"`   // 文章来源
	Link     string `json:"link"`     //原文连接

	BannerID uint        `json:"banner_id"` //文章封面ID
	Tags     ctype.Array `json:"tags"`      //文章标签
}

func (ArticleApi) ArticleCreateView(c *gin.Context) {
	var cr ArticleRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID
	userNickName := claims.NickName
	//校验context xss攻击

	//处理content
	unsafe := blackfriday.MarkdownCommon([]byte(cr.Content))
	//是不是有script标签
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	nodes := doc.Find("script").Nodes
	if len(nodes) > 0 {
		//有script标签，不允许发布
		doc.Find("script").Remove()
		converter := md.NewConverter("", true, nil)
		html, _ := doc.Html()
		markdown, _ := converter.ConvertString(html)
		cr.Content = markdown
	}
	if cr.Abstact == "" {
		//汉字的截取不一样
		abs := []rune(doc.Text())
		if len(abs) > 100 {
			cr.Abstact = string(abs[:100])
		} else {
			cr.Abstact = string(abs[:])
		}

	}

	// 不传bannerid
	if cr.BannerID == 0 {
		var bannerIdList []uint
		global.DB.Model(models.BannerModel{}).Select("id").Scan(&bannerIdList)
		if len(bannerIdList) == 0 {
			res.FailWithMessage("没有可用的banner", c)
		}
		rand.Seed(time.Now().UnixNano())
		cr.BannerID = bannerIdList[rand.Intn(len(bannerIdList))]
	}

	//查banner_id下的banner_url
	var bannerUrl string
	err = global.DB.Model(models.BannerModel{}).Where("id =?", cr.BannerID).Select("path").Scan(&bannerUrl).Error
	if err != nil {
		res.FailWithMessage("banner不存在", c)
		return
	}

	//查用户头像
	var avatar string
	err = global.DB.Model(models.UserModel{}).Where("id =?", userID).Select("avatar").Scan(&avatar).Error
	if err != nil {
		res.FailWithMessage("banner不存在", c)
		return
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	article := models.ArticleModel{
		CreateAt:     now,
		Update:       now,
		Title:        cr.Title,
		Keyword:      cr.Title,
		Abstact:      cr.Abstact,
		Content:      cr.Content,
		UserID:       userID,
		UserNickName: userNickName,
		UserAvatar:   avatar,
		Category:     cr.Category,
		Source:       cr.Source,
		Link:         cr.Link,
		BannerID:     cr.BannerID,
		BannerUrl:    bannerUrl,
		Tags:         cr.Tags,
	}
	if article.ISExistData() {
		res.FailWithMessage("文章已经存在", c)
		return
	}
	err = article.Create()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	go es_ser.AsyncArticleByFullText(article.ID, article.Title, article.Content)
	res.OKWithMessage("文章发布成功", c)

}
