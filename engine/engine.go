package engine

import (
	"crawler/fetcher"
	"log"
)

// Run把需要爬的网页url及解析器放进来
// seeds可以是任意多个Request结构的参数
func Run(seeds ...Request) {
	// 把seeds放进requests切片
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	// 遍历requests
	for len(requests) > 0 {
		// 读取一个request并左移切片
		r := requests[0]
		requests = requests[1:]

		// 用fetcher读取url内容
		log.Printf("Fetching %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
			continue
		}

		// 用parser解析fetcher读取的网页内容
		parseResult := r.ParserFunc(body)
		// 追加engine结下来需要处理的requests
		requests = append(requests, parseResult.Requests...)

		// 处理parser解析出来的item，也就是爬到的有用的数据
		for _, item := range parseResult.Items {
			log.Printf("Got Item %v", item)
		}
	}
}
