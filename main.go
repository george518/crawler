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
	"github.com/george518/crawler/persist"
	"github.com/george518/crawler/scheduler"
	"github.com/george518/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemServer(),
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
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
