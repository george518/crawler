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

	"github.com/george518/crawler/engine"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
)

func ItemServer(index string) (chan engine.Item, error) {

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("item saver: got item #%d:%v", itemCount, item)
			itemCount++
			err := Save(item, client, index)
			if err != nil {
				log.Printf("err #%d %v", itemCount, err)
			}
		}

	}()
	return out, nil
}

func Save(item engine.Item, client *elastic.Client, index string) error {

	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type)

	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.BodyJson(item).Do(context.Background())
	return err

}
