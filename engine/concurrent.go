/************************************************************
** @Description: engine
** @Author: haodaquan
** @Date:   2018-03-18 15:00
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-03-18 15:00
*************************************************************/
package engine

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan Item
	RequestProcessor Processor
}
type Processor func(request Request) (ParseResult, error)
type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	Run()
	ReadyNotifier
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	//e.Scheduler.ConfigureMasterWorkChan(in)
	//创建worker
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			//fmt.Printf("Got Item:%v\n", item)
			go func() { e.ItemChan <- item }()
		}

		for _, Request := range result.Requests {
			e.Scheduler.Submit(Request)
		}

	}
}

func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- result
		}

	}()
}
