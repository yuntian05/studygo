package engine

type Scheduler interface {
	ReadyNotify
	Submit(request Request)
	WorkChan() chan Request
	Run()
}

type ReadyNotify interface {
	WorkerReady(chan Request)
}

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan Item
	RequestProcessor Processor
}

type Processor func(r Request) (ParseResult, error)

func (e *ConcurrentEngine) Run(seeds ...Request)  {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i:= 0; i < e.WorkerCount; i++{
		e.createWorker(e.Scheduler, e.Scheduler.WorkChan(), out)
	}

	for _,r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <- out
		for _, item := range result.Items {
			go func() {e.ItemChan <- item}()
			//log.Printf("Got item:%v", item)
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func (e *ConcurrentEngine)createWorker(ready ReadyNotify, in chan Request,  out chan ParseResult) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <- in
			result, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visited = make(map[string]bool)

func isDuplicate(url string) bool  {
	if visited[url] {
		return true
	}
	visited[url] = true
	return false
}
