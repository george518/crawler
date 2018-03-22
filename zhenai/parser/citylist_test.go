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

	iconvString, err := iconv.ConvertString(string(contents),
		"GB2312",
		"utf-8")
	result := ParseCityList([]byte(iconvString), "")
	const resultSize = 470

	if len(result.Requests) != resultSize {
		t.Errorf("result shoud have %d,"+
			"requests,but had %d", resultSize, len(result.Requests))
	}

}
