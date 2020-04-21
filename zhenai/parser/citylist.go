package parser

import (
	"crawler/engine"
	"regexp"
)

// 用于匹配需要提取url和item的正则表达式
const cityListRe = `<a [^ ]*href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

// 解析城市列表的parser
func ParseCityList(contents []byte) engine.ParseResult {

	// 获取正则表达式匹配到的内容
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	// 初始化一个空的ParseResult
	result := engine.ParseResult{}

	// 把匹配到的内容逐个放进ParseResult
	for _, m := range matches {
		result.Items = append(
			result.Items, string(m[2]))
		result.Requests = append(
			result.Requests, engine.Request{
				Url:        string(m[1]),
				ParserFunc: engine.NilParser, //当还没有决定request用什么parserfunc的时候，使用NilParser
			})
	}

	return result
}
