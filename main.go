// http://cheez-2.local:57104/Dash/qwbydnau/godoc.org/golang.org/x/tools/benchmark/parse.html#Benchmark
package main

import (
	"bytes"
	"sync"
)

func main(){
	mut()
}

func mut(){
	var b bytes.Buffer
	var wg sync.WaitGroup
	m := new(sync.Mutex)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(){
			defer func(){
				m.Unlock()
				wg.Done()
			}()
			m.Lock()
			b.Write([]byte("Goose and gopher")) // Writing 16B
		}()
	}
	wg.Wait()
}
