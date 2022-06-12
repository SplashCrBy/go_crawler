package engine

import (
	"log"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		ParseResult, err := worker(r)
		if err != nil {
			continue
		}

		requests = append(requests, ParseResult.Requests...)

		for _, item := range ParseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}
