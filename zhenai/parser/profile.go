package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)

// "basicInfo": ["离异", "45岁", "魔羯座(12.22-01.19)", "177cm", "75kg", "工作地:广州天河区", "月收入:1.2-2万", "总监", "大专"],"detailInfo":["汉族","籍贯:广东广州","体型:壮实","社交场合会抽烟","稍微喝一点酒","租房","未买车","有孩子且偶尔会一起住","是否想要孩子:视情况而定","何时结婚:时机成熟就结婚"],"educationString":"大专","emotionStatus":0,"gender":0,"genderString":"男士"
const basicRe = `"basicInfo":[^"]*"([^"]+)"[^"]*"([\d]+)[^"]*"[^"]*"([^"]+)"[^"]*"([\d]+)[^"]*"[^"]*"([\d]+)[^"]*"[^"]*"([^"]+)"[^"]*"([^"]+)"[^"]*"([^"]+)"[^"]*"([^"]+)"[^:]+:[^"]*"([^"]+)"[^:]+:([^"]+)"[^:]+:([^"]+)"[^"]*"([^"]+)"[^"]*"([^"]+)"[^"]*"([^"]+)"[^"]*"([^"]+)"[^"]*"([^"]+)"[^:]+:([^"]+)"[^:]+:([^"]+)"[^:]+:[^"]*"([^"]+)"[^:]+:([^,]+),[^:]+:([^,]+),[^:]+:[^"]*"([^"]+)"`

// 实现Request所需ParseFunc，闭包增加一个name参数，在调用时传入上层ParseCity得到的name
func ParseProfile(contents []byte, name string) engine.ParseResult {

	profile := model.Profile{}
	profile.Name = name

	// 获取正则表达式匹配到的内容
	re := regexp.MustCompile(basicRe)
	match := re.FindSubmatch(contents)

	// 匹配到的内容放进model.Profile
	if match != nil {

		profile.Age, _ = strconv.Atoi(string(match[2]))
		profile.Height, _ = strconv.Atoi(string(match[4]))
		profile.Weight, _ = strconv.Atoi(string(match[5]))
		profile.Gender = string(match[23])
		profile.Income = string(match[7])
		profile.Marriage = string(match[1])
		profile.Education = string(match[9])
		profile.Occupation = string(match[8])
		profile.Hokou = string(match[11])
		profile.Xinzuo = string(match[3])
		profile.House = string(match[15])
		profile.Car = string(match[16])

	}

	return engine.ParseResult{
		Items: []interface{}{profile},
	}
}
