package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/reactivego/multicast"
)

func main() {
	type doin_something func(value interface{}, err error, closed bool) bool
	ipc := multicast.NewChan(110, 10)
	// ipc.NewEndpoint()
	// ipc.Close()
	wg := &sync.WaitGroup{}
	wg.Add(10)
	for i := range 10 {
		go func(i int) {
			new_chan, err := ipc.NewEndpoint(multicast.ReplayAll)
			if err != nil {
				return
			}
			work := func(i int) doin_something {
				log.Println("work for ", i, " has starteds")
				return func(value interface{}, err error, closed bool) bool {
					switch {
					case !closed:
						fmt.Printf("%v at %v\n", value, i)
					case err != nil:
						fmt.Printf("%v at %v\n", err, i)
					default:
						fmt.Printf("closed at %v \n", i)
					}
					return true
				}
			}
			new_chan.Range(
				work(i),
				0)
			wg.Done()
		}(i)
	}
	for i := range 100 {
		ipc.Send(i)
	}

	ipc.Close(nil)
	wg.Wait()
}
