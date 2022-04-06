package main

import (
	"fmt"
	"sync"
)
var wg = sync.WaitGroup{}
var m = sync.Mutex{};
func main() {
	writer := ConsoleWriter{};
	for i := 0; i < 10; i++ {
		fmt.Println()
		wg.Add(3)
		m.Lock()
		go func(writer ConsoleWriter){
			writer.WideWrite([]byte("Hello, world!"))
			m.Unlock()
			wg.Done()
		}(writer)
		m.Lock()
		go func(writer ConsoleWriter) {
			writer.WideWrite([]byte("I am near your house."))
			m.Unlock()
			wg.Done()
		}(writer) 
		fmt.Println()
		m.Lock()
		go func(writer ConsoleWriter) {
			writer.WideWrite([]byte("YOU MUST EVACUATE THE UNIVERSE."))
			m.Unlock()
			wg.Done()
		}(writer)
		fmt.Println()
	}
	wg.Wait()
}
type Writer interface {
	Write([]byte) (int, error)
}
type ConsoleWriter struct {
}

func (cw ConsoleWriter) WideWrite(b []byte) (int, error)  {
	var err error
	var n int
	for i := 0; i < len(b); i++ {
		n, err = fmt.Print(string(b[i]), "  ")
	}
	return n, err
}