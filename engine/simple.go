/************************************************************
** @Description: engine
** @Author: haodaquan
** @Date:   2018-03-17 16:36
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-03-17 16:36
*************************************************************/
package engine

import (
	"github.com/george518/crawler/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, err := Worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("got item %v", item)
		}
	}
}

func Worker(r Request) (ParseResult, error) {
	//log.Printf("fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		//log.Printf("fetch: error fetch url%s:%v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body), nil
}
