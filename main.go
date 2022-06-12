package main

import (
	"crawlers/engine"
	"crawlers/persist"
	"crawlers/scheduler"
	"crawlers/zhenai/parser"
)

const url = "http://localhost:8080/mock/www.zhenai.com/zhenghun"

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	e.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
	})

}
