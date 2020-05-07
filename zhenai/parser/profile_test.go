package parser

import (
	"crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	//读取测试用例页面文件
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%s\n", contents)

	//获取所测试的ParseProfile结果
	result := ParseProfile(contents, "拐角遇到你")

	//测试返回长度测试
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 "+"element; but was %v", result.Items)
	}

	//获取测试结果的Item项内容
	profile := result.Items[0].(model.Profile)

	//所需对比的期望内容
	expected := model.Profile{
		Name:       "拐角遇到你",
		Age:        45,
		Height:     177,
		Weight:     75,
		Gender:     "男士",
		Income:     "月收入:1.2-2万",
		Marriage:   "离异",
		Education:  "大专",
		Occupation: "总监",
		Hokou:      "广东广州",
		Xinzuo:     "魔羯座(12.22-01.19)",
		House:      "租房",
		Car:        "未买车",
	}

	//对比测试内容
	if profile != expected {
		t.Errorf("expected %v: but was %v", expected, profile)
	}

}
