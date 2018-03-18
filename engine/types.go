/************************************************************
** @Description: engine
** @Author: haodaquan
** @Date:   2018-03-17 16:11
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-03-17 16:11
*************************************************************/
package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
