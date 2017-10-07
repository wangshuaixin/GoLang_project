package main

// 基础包
import (
	"net/http"
	"strings"

	"github.com/henrylee2cn/pholcus/app/downloader/request"
	"github.com/henrylee2cn/pholcus/common/goquery"
	// "github.com/henrylee2cn/pholcus/common/goquery"  //DOM解析
	// "github.com/henrylee2cn/pholcus/app/downloader/request" //必需
	. "github.com/henrylee2cn/pholcus/app/spider" //必需
	// . "github.com/henrylee2cn/pholcus/app/spider/common" //选用
	// "github.com/henrylee2cn/pholcus/logs"
	// net包
	//  "net/http" //设置http.Header
	// "net/url"
	// 编码包
	// "encoding/xml"
	//"encoding/json"
	// 字符串处理包
	//"regexp"
	// "strconv"
	// "strings"
	// 其他包
	"fmt"
	// "math"
	// "time"
)

func inti() {
	FileTest1.Register()
}

var FileTest1 = &Spider{
	Name:        "Pic Downloader", //url:http//meizitu.com/
	Description: "Pic Downloader",
	//Pausetime: 300,
	//Keyin: KEYIN,
	//Limit: LIMIT,
	EnableCookie: false,
	RuleTree: &RuleTree{
		Root: func(ctx *Context) {
			ctx.AddQueue(&request.Request{
				Url:          "http://www.,meizitu.com/",
				Rule:         "meizitu",
				ConnTimeout:  -1,
				DownloaderID: 0, // pic or other media must use 0; surfer surf raw downloader
			})
		},
		Trunk: map[string]*Rule{
			"meizitu": {
				ParseFunc: func(ctx *Context) {
					query := ctx.GetDom()
					query.Find("#picture p a").Each(func(i int, s *goquery.Selection) {
						//其中picture是id，
						//p是tagname a也是tagname，这里就找到了a，而且是好多个，所以用了each进行循环
						//s.Find('p>a')
						fmt.Println("Print test.")
						fmt.Println(s.Html())
						fmt.Println("--------------------")
						url1, _ := s.Attr("href")
						fmt.Println(url1)
						fmt.Println("?????????????????????")
						//t:=s.Find('a').Eq(0)
						//fmt.Println("A的html",t.Html())
						if href, ok := s.Attr("href"); ok {
							ctx.AddQueue(&request.Request{
								Url:    href,
								Header: http.Header{"Content-Type": []string{"text/html; charset=gbk"}},
								Rule:   "图片URL",
							})
						}
					})
				},
			},
			"Pic URL": {
				ParseFunc: func(ctx *Context) {
					query := ctx.GetDom()
					query.Find("#picture p img").Each(func(i int, s *goquery.Selection) {
						fmt.Println("Picture test2")
						fmt.Println(s.Html())
						fmt.Println("------------------")
						url1, _ := s.Attr("src")
						fmt.Println(url1)
						fmt.Println("?????????????????")
						fmt.Println("End this secondary pages.")
						//t:=s.Find('a').Eq(0)
						//fmt.Println("A的html",t.Html())
						if href, ok := s.Attr("src"); ok {
							ctx.AddQueue(&request.Request{
								Url:    href,
								Header: http.Header{"Content-Type": []string{"text/html; charset=gbk"}},
								Rule:   "图片下载",
							})
						}
					})
				},
			},
			"Pic Download": {
				ParseFunc: func(ctx *Context) {
					fmt.Println("Picture URL: ", ctx.GetUrl())
					picurl := ctx.GetUrl()                           // get the url, use named with picture; may be related to mkdir or agr
					picname = strings.Replace(picname, "/", "-", -1) //replace value
					ctx.FileOutput(picname)                          // equal to ctx.AddFile()
				},
			},
		},
	},
}
