package gpool

type Executor struct {
	pool	*Pool
	taskC	chan func()
	stop	chan struct{}
}

type Pool struct {
	es			chan *Executor
	taskList	chan func()
	stop		chan struct{}
}

func (e *Executor) Start() {
	go func() {
		var t func()
		for {
			e.pool.es <- e
			select {
			case <-e.stop:
				return
			case t = <-e.taskC:
				t()
			}
		}
	}()
}

func NewGPool(maxInvoke, queueCap int) *Pool {
	p := &Pool{
		es:			make(chan *Executor, maxInvoke),
		taskList:	make(chan func(),queueCap),
		stop:		make(chan struct{}),
	}
	p.Start()
	return p
}

func (p *Pool) Start() {
	for i := 0; i < cap(p.es); i++ {
		e := Executor{
			pool:	p,
			taskC:	make(chan func()),
			stop:	make(chan struct{}),
		}
		e.Start()
	}
	go func() {
		for {
			select {
			case (<-p.es).taskC <- (<-p.taskList):
			case <-p.stop:
				for i := 0; i < cap(p.es); i++ {
					(<-p.es).stop <- struct{}{}
				}
				return
			}
		}
	}()
}

func (p *Pool) Release() {
	p.stop <- struct{}{}
}

func (p *Pool) AddTask(t func()) {
	p.taskList <- t
}