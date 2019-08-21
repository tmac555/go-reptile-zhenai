package schedular

import "reptile/enginer"

type Simpleschedular struct {
	Workchan chan enginer.Request
}

func (s *Simpleschedular) Masterchan(w chan enginer.Request) {
	s.Workchan = w
}
func (s *Simpleschedular) Submit(r enginer.Request) {
	go func() { s.Workchan <- r }()
}
