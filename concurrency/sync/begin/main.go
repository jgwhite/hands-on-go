// concurrency/sync/begin/main.go
package main

import (
	"fmt"
	"sync"
)

func main() {
	// given a list of names
	names := []string{"James", "Toni", "Marcus"}

	// initialize a map to store the length of each name
	nameMap := make(map[string]int)
	var nameMapMu sync.Mutex

	// initialize a wait group for the goroutines that will be launched
	var wg sync.WaitGroup
	wg.Add(len(names))

	// launch a goroutine for each name we want to process
	for _, name := range names {
		go func(name string) {
			defer wg.Done()
			nameMapMu.Lock()
			defer nameMapMu.Unlock()
			nameMap[name] = len(name)
		}(name)
	}

	// wait for all goroutines to finish
	wg.Wait()

	// print the map
	fmt.Println(nameMap)
}
