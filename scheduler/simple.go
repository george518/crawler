/************************************************************
** @Description: scheduler
** @Author: haodaquan
** @Date:   2018-03-18 16:00
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-03-18 16:00
*************************************************************/
package scheduler

import (
	"github.com/george518/crawler/engine"
)

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {

}

func (s *SimpleScheduler) Submit(
	r engine.Request) {
	go func() { s.workerChan <- r }()
}
