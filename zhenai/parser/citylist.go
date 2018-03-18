/************************************************************
** @Description: parser
** @Author: haodaquan
** @Date:   2018-03-17 16:09
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-03-17 16:09
*************************************************************/
package parser

import (
	"crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(content []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	mth := re.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
	limit := 100
	for _, m := range mth {
		//result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})

		limit--
		if limit == 0 {
			break
		}
	}

	return result
}
