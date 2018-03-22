/************************************************************
** @Description: engine
** @Author: haodaquan
** @Date:   2018-03-21 23:27
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-03-21 23:27
*************************************************************/
package engine

import "github.com/george518/crawler/fetcher"

func Worker(r Request) (ParseResult, error) {
	//log.Printf("fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		//log.Printf("fetch: error fetch url%s:%v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body, r.Url), nil
}
