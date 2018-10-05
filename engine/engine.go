package engine

import (
	"depthLearn/goCrawler/fetcher"
	"fmt"
	"log"
)

func Run(seed ...Request) {
	var requests []Request
	for _, r := range seed {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, err := worker(r)
		if err != nil{
			continue
		}
		requests = append(requests,
			parseResult.Requests...)

		for _, item := range parseResult.Items {
			fmt.Println("Get item %v", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	//log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error " + "fetching url %s: %v", r.Url, err)
		return ParseResult{}, nil
	}

	return r.ParserFunc(body), nil
}