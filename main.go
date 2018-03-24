/************************************************************
** @Description: crawler
** @Author: haodaquan
** @Date:   2018-03-17 14:23
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-03-17 14:23
*************************************************************/
package main

import (
	"github.com/george518/crawler/engine"
	"github.com/george518/crawler/scheduler"
	"github.com/george518/crawler/zhenai/parser"
	"github.com/george518/crawler_distributed/persist/client"
)

func main() {
	//itemChan, err := persist.ItemServer("dating_profile")
	itemChan, err := client.ItemServer(":1234")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}
	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})

	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/xuzhou",
	//	ParserFunc: parser.ParseCity,
	//})

	//e := engine.ConcurrentEngine{
	//	Scheduler:   &scheduler.SimpleScheduler{},
	//	WorkerCount: 100,
	//}
	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

}
