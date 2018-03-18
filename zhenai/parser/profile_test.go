/************************************************************
** @Description: parser
** @Author: haodaquan
** @Date:   2018-03-17 23:15
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-03-17 23:15
*************************************************************/
package parser

import (
	"io/ioutil"
	"testing"

	"fmt"

	iconv "github.com/djimenez/iconv-go"
)

func TestParseProfile(t *testing.T) {
	content, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	iconvString, err := iconv.ConvertString(string(content), "GB2312", "utf-8")
	result := ParseProfile([]byte(iconvString), "")

	if len(result.Items) != 1 {
		t.Errorf("result shoud contain 1 element,but was %v", result.Items)
	}

	fmt.Println(result.Items)
}
