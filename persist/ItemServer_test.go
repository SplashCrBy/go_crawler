package persist

import (
	"context"
	"crawlers/engine"
	"crawlers/model"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"testing"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url: "http://localhost:8080/mock/album.zhenai.com/u/8256018539338750764",
		Id:  "8256018539338750764",
		Payload: model.Profile{
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
		},
	}
	const index = "dating_test"

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	err = save(client, index, expected)
	if err != nil {
		panic(err)
	}
	resp, err := client.Get().Index(index).Id(expected.Id).Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

	var actual engine.Item
	err = json.Unmarshal(resp.Source, &actual)

	if err != nil {
		panic(err)
	}

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile
	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
