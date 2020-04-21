package engine

// 用来接收需要爬取的页面url和相应的解析函数
type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

// 用来返回解析后得到
//        1. 新页面请求列表(切片)
//        2. 所需数据列表(切片)
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

// 临时用于在解析器里返回一个空的解析结果
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
