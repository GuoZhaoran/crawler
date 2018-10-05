package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler Scheduler         //调度器
	WorkerCount int             //配置worker的数量
}

type Scheduler interface {
	Submit(Request)
	ConfigureWorkerMasterChan(chan chan Request)
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0;i < e.WorkerCount;i++ {
       createWorker(out, e.Scheduler)
	}
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	for {
		result := <- out
		for _,item := range result.Items{
			fmt.Println(item)
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(out chan ParseResult,s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}