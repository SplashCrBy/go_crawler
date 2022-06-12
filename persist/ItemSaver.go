package persist

import (
	"context"
	"crawlers/engine"
	"github.com/olivere/elastic/v7"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	out := make(chan engine.Item)
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("item saver  : got item #%d: %v", itemCount, item)
			itemCount++

			err := save(client, index, item)
			if err != nil {
				log.Printf("error: %s", err)
			}
		}

	}()
	return out, nil
}

func save(client *elastic.Client, index string, item engine.Item) error {
	indexService := client.Index().Index(index).OpType("create").Id(item.Id).BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}
