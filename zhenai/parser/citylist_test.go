/************************************************************
** @Description: parser
** @Author: haodaquan
** @Date:   2018-03-17 16:53
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-03-17 16:53
*************************************************************/
package parser

import (
	"io/ioutil"
	"testing"

	"github.com/djimenez/iconv-go"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data_utf8.html")
	if err != nil {
		panic(err)
	}

	iconvString, err := iconv.ConvertString(string(contents), "GB2312", "utf-8")

	//fmt.Println(iconvString)

	result := ParseCityList([]byte(iconvString))
	const resultSize = 470

	if len(result.Requests) != resultSize {
		t.Errorf("result shoud have %d,"+
			"requests,but had %d", resultSize, len(result.Requests))
	}

	if len(result.Items) != resultSize {
		t.Errorf("result shoud have %d,"+
			"requests,but had %d", resultSize, len(result.Items))
	}

	expectedCities := []string{
		"阿坝",
		"阿克苏",
		"阿拉善盟",
	}

	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d:%s but was %s",
				i, url, result.Requests[i].Url)
		}
	}

	for i, item := range expectedCities {
		if result.Items[i].(string) != item {
			t.Errorf("expected city #%d:%s but was %s",
				i, item, result.Items[i].(string))
		}
	}

}
