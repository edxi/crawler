package parser

import (
	"crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {

	// 获取正则表达式匹配到的内容
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	// 初始化一个空的ParseResult
	result := engine.ParseResult{}

	// 把匹配到的内容逐个放进ParseResult
	for _, m := range matches {
		result.Items = append(
			result.Items, "User "+string(m[2]))
		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(m[1]),
				// 函数式编程实现闭包调用
				ParserFunc: func(c []byte) engine.ParseResult {
					return ParseProfile(c, string(m[2]))
				},
			})
	}

	return result

}
