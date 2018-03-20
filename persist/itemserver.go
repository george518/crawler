/************************************************************
** @Description: persist
** @Author: haodaquan
** @Date:   2018-03-18 21:52
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-03-18 21:52
*************************************************************/
package persist

import (
	"log"

	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
)

func ItemServer() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			//log.Printf("item saver: got item #%d %v", itemCount, item)
			itemCount++
			id, err := save(item)
			if err != nil {
				log.Printf("err #%d %v", itemCount, err)
			}

			log.Printf("ok #%s", id)
		}

	}()
	return out
}

func save(item interface{}) (string, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return "", err
	}

	resp, err := client.Index().
		Index("data_profile").
		Type("zhenai").
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		return "", err
	}

	return resp.Id, nil

}
