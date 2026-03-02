package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m       = sync.RWMutex{}
	wg      = sync.WaitGroup{}
	dbData  = []string{"id1", "id2", "id3", "id4", "id5"}
	results = []string{}
)

func main() {
	t0 := time.Now()
	for i := range len(dbData) {
		wg.Add(1)
		go dbCall(i)
	}

	wg.Wait()
	fmt.Printf("\nTotal execution time: %v\n", time.Since(t0))
	log()
}

func dbCall(i int) {
	delay := 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is: ", dbData[i])
	save(dbData[i])
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("\nthe current results are: %v", results)
	m.RUnlock()
}
