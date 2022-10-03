package crawler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
	"log"
	"strings"
	"time"
)

func Search(kw string) ([][]string, error) {

	datas := [][]string{}
	url := fmt.Sprintf(`https://zh.wikihow.com/wikiHowTo?search=%s`, kw)
	s := gorequest.New().Timeout(12 * time.Second)
	resp, _, err := s.Get(url).
		Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 Edge/16.16299").
		End()

	if err != nil {
		log.Printf("get wikihow search url error:%v\\n", err)
		return datas, err[0]
	}

	if resp.StatusCode == 404 {
		return datas, fmt.Errorf("404 Page Not Found!")
	}

	doc, err_ := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("goquery NewDocumentFromReader error:%v\\n", err_)
		return datas, err_
	}

	//aNodes := doc.Find(`div#searchresults_list>a.result_link`).Nodes
	//for idx, a := range aNodes {
	//	href := ""
	//	if idx == 0 {
	//		attrs := a.Attr
	//
	//		for _, attr := range attrs {
	//			if attr.Key == "href" {
	//				href = attr.Val
	//				break
	//			}
	//		}
	//
	//		if href == "https://zh.wikihow.com/%E9%A6%96%E9%A1%B5" {
	//			break
	//		}
	//	}
	//
	//}

	doc.Find(`div#searchresults_list>a.result_link`).Each(func(i int, a *goquery.Selection) {
		href, _ := a.Attr("href")

		title := a.Find("div.result_title").Text()

		view := a.Find(`li.sr_view`).Text()
		data := []string{}
		for _, d := range []string{title, view, href} {
			val := strings.ReplaceAll(d, " ", "")
			val = strings.ReplaceAll(val, "\n", "")
			val = strings.ReplaceAll(val, "\r", "")
			val = strings.ReplaceAll(val, "\t", "")
			data = append(data, val)
		}
		datas = append(datas, data)
	})

	return datas, nil
}
