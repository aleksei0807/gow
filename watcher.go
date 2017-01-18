package main

import (
	"log"

	"sync"

	fswatch "gopkg.in/andreaskoch/go-fswatch.v1"
)

const (
	checkIntervalInSeconds = 1
)

func watcher(file string, initParams params, wg *sync.WaitGroup) {
	fileWatcher := fswatch.NewFileWatcher(file, checkIntervalInSeconds)

	nullChange := false

	fileWatcher.Start()

	for fileWatcher.IsRunning() {

		select {
		case <-fileWatcher.Modified():
			go func() {
				if nullChange {
					log.Printf(`file "%s" is changed`, file)
				}
				nullChange = true
			}()

		case <-fileWatcher.Moved():
			go func() {
				log.Printf(`file "%s" is moved`, file)
				fileWatcher.Stop()
				wg.Done()
				log.Println("done!")

				watchPath(initParams, wg, true)
			}()
		}

	}
}
