package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
}

/*
type Scheduler interface {
	Submit(request Request)
	ConfigureMasterWorkerChan(chan Request)
}
*/
type Scheduler interface {
	ReadyNotifier
	Submit(request Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	//in := make(chan Request)
	out := make(chan ParseResult)
	//e.Scheduler.ConfigureMasterWorkerChan(in)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		//createWorker(in, out)
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			//log.Printf("Duplicate request: %s", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}
	for {
		result := <-out

		for _, item := range result.Items {
			go func(i Item) { e.ItemChan <- i }(item)

		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				//log.Printf("Duplicate request: %s", request.Url)
				continue
			}
			e.Scheduler.Submit(request)
		}
	}

}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

/*
func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}*/

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
