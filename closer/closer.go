package closer

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
)

type Closer struct {
	mu    sync.Mutex
	funcs []Func
}

type Func func() error

func (c *Closer) Add(cls ...Func) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.funcs = append(c.funcs, cls...)
}

func New(signals ...os.Signal) *Closer {
	cl := new(Closer)
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, signals...)
		<-ch
		signal.Stop(ch)
		cl.Close()
	}()
	return cl
}

func (c *Closer) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for i := len(c.funcs) - 1; i >= 0; i-- {
		item := c.funcs[i]
		if err := item(); err != nil {
			log.Println(err.Error())
		}
	}
	fmt.Println("Gracefully shutting down")
}
