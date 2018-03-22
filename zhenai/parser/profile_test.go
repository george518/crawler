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

	"regexp"

	iconv "github.com/djimenez/iconv-go"
)

func TestParseProfile(t *testing.T) {
	content, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	iconvString, err := iconv.ConvertString(string(content), "GB2312", "utf-8")
	result := ParseProfile([]byte(iconvString), "", "")

	if len(result.Items) != 1 {
		t.Errorf("result shoud contain 1 element,but was %v", result.Items)
	}

	fmt.Println(result.Items)
}

func TestExtractStringr(t *testing.T) {
	url := "http://album.zhenai.com/u/108297578"
	var IdUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)
	NO := extractString([]byte(url), IdUrlRe)

	fmt.Println(NO)
}
