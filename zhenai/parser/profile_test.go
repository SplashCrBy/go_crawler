package parser

import (
	"crawlers/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "寂寞成影萌宝")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)
	expected := model.Profile{
		Name:       "寂寞成影萌宝",
		Gender:     "女",
		Age:        83,
		Height:     105,
		Weight:     137,
		Income:     "财务自由",
		Marriage:   "离异",
		Education:  "初中",
		Occupation: "金融",
		Hukou:      "南京市",
		Xinzuo:     "狮子座",
		House:      "无房",
		Car:        "无车",
	}

	if profile != expected {
		t.Errorf("expected %v; but was %v", expected, profile)
	}
}
