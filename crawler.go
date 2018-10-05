package main

import (
	"depthLearn/goCrawler/engine"
	"depthLearn/goCrawler/samecity/parser"
	"depthLearn/goCrawler/scheduler"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount:100,
	}
	e.Run(engine.Request{
		Url:"https://bj.58.com/chuzu/?PGTID=0d100000-0000-101e-c465-298d85d88a10&ClickID=8",
		ParserFunc:parser.ParseCity,
	})
}

