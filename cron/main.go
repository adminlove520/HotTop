package main

import (
	"crawler/lib"
	"github.com/robfig/cron"
	"fmt"
	"log"
)

var linkPool = make(chan lib.Link, 3)
var pagePool = make(chan lib.Page, 3)

func AddSite(s lib.Spider) {
	links, _ := s.BuildUrl()
	for _, link := range links {
		go func(link lib.Link) {
			linkPool <- link
		}(link)
	}
}

func AddSites() {
	var spList []lib.Spider
	spList = append(spList, &lib.V2ex{
		Site: lib.Site{
			Root: "https://www.v2ex.com",
			Domain: "www.v2ex.com",
			Desc: "way to explore",
			CrawType:lib.CrawHtml,
		},
	})
	spList = append(spList, &lib.Chouti{
		Site: lib.Site{
			Root: "https://dig.chouti.com",
			Domain: "https://dig.chouti.com",
			Desc: "抽屉新热榜",
			CrawType:lib.CrawApi,
		},
	})
	spList = append(spList, &lib.Zhihu{
		Site: lib.Site{
			Root: "https://zhihu.com",
			Domain: "www.zhihu.com",
			Desc: "知乎热榜",
			CrawType:lib.CrawHtml,
		},
	})
	spList = append(spList, &lib.Weibo{
		Site: lib.Site{
			Root: "https://s.weibo.com",
			Domain: "www.weibo.com",
			Desc: "微博热搜",
			CrawType:lib.CrawHtml,
		},
	})

	for _, v := range spList {
		go AddSite(v)
	}
}

func CrawSite() {
	for {
		select {
		case l := <- linkPool:
			go func() {
				sp := l.Sp
				page, err := sp.CrawPage(l)
				if err != nil {
					fmt.Println("%w", err)
				}
				pagePool <- page
			}()
		case p := <- pagePool:
			go func() {
				sp := p.Link.Sp
				sp.Store(p)
			}()
		}
	}
}

func main() {
	cron := cron.New()
	log.Fatal(cron.AddFunc("0 * * * *", func() {
		fmt.Println("start crawling ...")
		AddSites()
	}))
	cron.Start()

	CrawSite()
}
