package db

import (
	"blogger/model"
	"testing"
	"time"
)

func init() {
	dns := "root:123@tcp(localhost:3306)/blogger?parseTime=true"
	err := InitDb(dns)
	if err != nil {
		panic(err)
	}
}

func testInsertArticle(t *testing.T) {
	articleDetail := model.ArticleDetail{}
	articleDetail.Category.CategoryId = 1
	articleDetail.Title = "我的第二篇文章，GOGOGOGOGOGGOlang"
	articleDetail.ViewCount = 0
	articleDetail.CommentCount = 0
	articleDetail.Username = "zzs"
	articleDetail.Summary = "Go是一种新的语言，一种并发的、带垃圾回收的、快速编译的语言。它具有以下特点：它可以在一台计算机上用几秒钟的时间编译一个大型的Go程序。"
	articleDetail.Content = `Go语言是一种新的语言，一种并发的、带垃圾回收的、快速编译的语言。它具有以下特点：
	1.它可以在一台计算机上用几秒钟的时间编译一个大型的Go程序。
	2.Go语言为软件构造提供了一种模型，它使依赖分析更加容易，且避免了大部分C风格include文件与库的开头。
	3.Go语言是静态类型的语言，它的类型系统没有层级。因此用户不需要在定义类型之间的关系上花费时间，这样感觉起来比典型的面向对象语言更轻量级。
	4.Go语言完全是垃圾回收型的语言，并为并发执行与通信提供了基本的支持。`
	articleDetail.CreateTime = time.Now()
	articleId, err := InsertArticle(&articleDetail)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("articleId = %d \n", articleId)
}

func testGetArticleList(t *testing.T) {
	list, err := GetArticleList(1, 10)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("list length = %d \n", len(list))
}

func TestGetArticleDetail(t *testing.T) {
	detail, _ := GetArticleDetail(1)

	t.Logf("detail：%v \n", detail.Id)
}

func TestGetPreArticleByArticleId(t *testing.T) {
	info, _ := GetNextArticleByArticleId(2)
	t.Logf("info: %v \n", info.ArticleID)
}
