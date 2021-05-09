package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	tr := NewTracker()
	go tr.Run()
	tr.Event(context.Background(), "test01")
	tr.Event(context.Background(), "test02")
	tr.Event(context.Background(), "test03")
	time.Sleep(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()
	tr.Shutdown(ctx)
}

//NewTracker new tracker
func NewTracker() *Tracker {
	return &Tracker{
		ch: make(chan string, 10),
	}
}

//Tracker type
type Tracker struct {
	ch   chan string
	stop chan struct{}
}

//Event jilu
func (t *Tracker) Event(ctx context.Context, data string) error {
	select {
	case t.ch <- data:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

//Run run
func (t *Tracker) Run() {
	for data := range t.ch {
		time.Sleep(time.Second)
		fmt.Println(data)
	}
	t.stop <- struct{}{}
}

//Shutdown close
func (t *Tracker) Shutdown(ctx context.Context) {
	close(t.ch)
	select {
	case <-t.stop:
		fmt.Println("t.stop coming")
	case <-ctx.Done():
		fmt.Println("ctx.Done coming")
	}
}
