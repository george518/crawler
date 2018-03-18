/************************************************************
** @Description: persist
** @Author: haodaquan
** @Date:   2018-03-18 21:52
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-03-18 21:52
*************************************************************/
package persist

import "log"

func ItemServer() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("item saver: got item #%d %v", itemCount, item)
			itemCount++
		}

	}()
	return out
}
