/************************************************************
** @Description: parser
** @Author: haodaquan
** @Date:   2018-03-17 21:55
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-03-17 21:55
*************************************************************/
package parser

import (
	"regexp"

	"github.com/george518/crawler/engine"
	"github.com/george518/crawler_distributed/config"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(` href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(content []byte, _ string) engine.ParseResult {
	mth := profileRe.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
	for _, m := range mth {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: NewProfileParser(string(m[2])),
		})
	}

	matches := cityUrlRe.FindAllSubmatch(content, -1)

	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
	}

	return result
}
