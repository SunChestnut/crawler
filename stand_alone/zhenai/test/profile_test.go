package test

import (
	"crawler/stand_alone/model"
	"crawler/stand_alone/zhenai/parser"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {

	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := parser.ParseProfile(contents, "萌宝")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)
	expected := model.Profile{
		Name:       "萌宝",
		Age:        40,
		Height:     176,
		Weight:     114,
		Income:     "3001-5000元",
		Gender:     "男",
		XinZuo:     "金牛座",
		Occupation: "金融",
		Marriage:   "离异",
		House:      "有房",
		HuKou:      "其它",
		Education:  "大学",
		Car:        "有豪车",
	}

	if expected != profile {
		t.Errorf("expected %v; but was %v", expected, profile)
	}

}
