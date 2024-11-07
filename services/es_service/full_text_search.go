package es_service

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/olivere/elastic/v7"
	"github.com/russross/blackfriday"
	"github.com/sirupsen/logrus"
	"server/global"
	"server/models"
	"strings"
)

type SearchData struct {
	Key   string `json:"key"`
	Body  string `json:"body"`  // 正文
	Slug  string `json:"slug"`  // 包含文章的id 的跳转地址
	Title string `json:"title"` // 标题
}

// AsyncArticleByFullText 同步文章数据到全文搜索
func AsyncArticleByFullText(id, title, content string) {
	indexList := GetSearchIndexDataByContent(id, title, content)

	// 创建一个桶
	bulk := global.EsClient.Bulk()
	for _, IndexData := range indexList {
		req := elastic.NewBulkIndexRequest().
			Index(models.FullTextModel{}.Index()).
			Doc(IndexData)
		bulk.Add(req)
	}
	result, err := bulk.Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Infof("%s 添加成功,共 %d 条", title, len(result.Succeeded()))
}

// DeleteFullTextByArticleID 删除文章数据
func DeleteFullTextByArticleID(id string) {
	// 用全文搜索去删
	boolSearch := elastic.NewTermQuery("key", id)
	res, _ := global.EsClient.
		DeleteByQuery().
		Index(models.FullTextModel{}.Index()).
		Query(boolSearch).
		Do(context.Background())
	logrus.Infof("成功删除 %d 条记录", res.Deleted)
}

func GetSearchIndexDataByContent(id, title, content string) (searchDataList []SearchData) {
	dataList := strings.Split(content, "\n") // 按照行去切割，以#为分割，分割成列表或切片
	var isCode bool = false                  // 默认不是代码块
	var headList, bodyList []string
	var body string
	headList = append(headList, getHeader(title))
	for _, s := range dataList {
		// #{1,6}
		// 判断一下是否是代码块
		if strings.HasPrefix(s, "```") { // 第一次遇见是代码块
			isCode = !isCode
		}
		if strings.HasPrefix(s, "#") && !isCode { // 第二次遇见就不是代码块
			headList = append(headList, getHeader(s))
			bodyList = append(bodyList, getBody(body))
			body = ""
			continue
		}
		body += s
	}
	bodyList = append(bodyList, getBody(body))
	ln := len(headList)
	for i := 0; i < ln; i++ {
		searchDataList = append(searchDataList, SearchData{
			Title: headList[i],
			Body:  bodyList[i],
			Slug:  id + getSlug(headList[i]),
			Key:   id,
		})
	}
	//b, _ := json.Marshal(searchDataList)
	//fmt.Println(string(b))
	return searchDataList
}

// getHeader 对文章标题进行处理
func getHeader(head string) string {
	head = strings.ReplaceAll(head, "#", "") // 把#替换成空
	head = strings.ReplaceAll(head, " ", "")
	return head
}

// getBody 对文章内容进行处理
func getBody(body string) string {
	unsafe := blackfriday.MarkdownCommon([]byte(body))
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	return doc.Text()
}

func getSlug(slug string) string {
	return "#" + slug
}
