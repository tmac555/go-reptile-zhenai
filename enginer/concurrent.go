package enginer

import (
	"fmt"
	"log"
	"reptile/fetch"
	"time"

	//"reptile/fetch"
)

type Concurrent struct {
	Schedular Schedular
	Workcount int
	Itemchan  chan Perfileitem
}

type Schedular interface {
	Submit(Request)
	Masterchan(chan Request)
}

//running
func (e *Concurrent) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Schedular.Masterchan(in)
	for i := 0; i <= e.Workcount; i++ {
		creatework(in, out)
	}
	for _, res := range seeds {
		e.Schedular.Submit(res)
	}

	for {
		result := <-out
		for _, request := range result.Request {
			e.Schedular.Submit(request)
		}
		for _, item := range result.Item {
			go func() { e.Itemchan <- item }()
		}
	}
	for {
		select {
		case <-time.After(5 * time.Second):
			fmt.Println("time out")
			return
		}
	}

}

//工作协程num
func creatework(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := Work(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()

}

//parseresult
func Work(r Request) (ParseResult, error) {
	body, err := fetch.Fetchrequest(r.Url)
	if (err != nil) {
		return ParseResult{}, err

	}
	log.Printf("fetch url=%s", r.Url)
	return r.Parsefunc(body, r.Url), nil
}
