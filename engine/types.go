/************************************************************
** @Description: engine
** @Author: haodaquan
** @Date:   2018-03-17 16:11
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-03-17 16:11
*************************************************************/
package engine

type ParserFunc func(content []byte, url string) ParseResult
type Request struct {
	Url        string
	ParserFunc ParserFunc
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Type    string
	Id      string
	PayLoad interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
