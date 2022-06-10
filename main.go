package main

import (
	"crawlers/engine"
	"crawlers/scheduler"
	"crawlers/zhenai/parser"
)

const url = "http://localhost:8080/mock/www.zhenai.com/zhenghun"

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}

	e.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
	})

}
