/************************************************************
** @Description: persist
** @Author: haodaquan
** @Date:   2018-03-20 23:07
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-03-20 23:07
*************************************************************/
package persist

import (
	"testing"

	"encoding/json"

	"github.com/george518/crawler/engine"
	"github.com/george518/crawler/model"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
)

func TestItemServer(t *testing.T) {
	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/108835456",
		Id:   "108835456",
		Type: "zhenai",
		PayLoad: model.Profile{
			Name:       "缱绻",
			Gender:     "女",
			Age:        28,
			Height:     165,
			Weight:     0,
			Income:     "3000元以下",
			Marriage:   "离异",
			Education:  "中专",
			Occupation: "--",
			Hukou:      "江苏徐州",
			Xingzuo:    "狮子座",
			House:      "--",
			Car:        "未购车",
			City:       "江苏徐州",
		},
	}
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	const index = "test_profile"

	err = Save(expected, client, index)

	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
		//t.Errorf(" error%v", err)
	}

	var actual engine.Item
	err = json.Unmarshal(*resp.Source, &actual)
	actualProfile, _ := model.FromJsonObj(actual.PayLoad)
	actual.PayLoad = actualProfile

	if expected != actual {
		t.Errorf("got %v,but expected %v", actual, expected)
	}
}
