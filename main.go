package main

// engine模块不断获取requet,交给不同的parser处理
import (
	"crawler/engine"
	"crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}
