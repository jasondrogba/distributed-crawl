package first_crawl_news

/*
从澎湃新闻的网站上获取首页的新闻内容。
*/
import (
	"fmt"
	"io"
	"net/http"
)

func PengpaiCrawl() {
	// http请求
	resp, err := http.Get("https://www.thepaper.cn/")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}
	// 获取返回的数据
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 打印返回的数据
	fmt.Println("body:", string(content))
}
