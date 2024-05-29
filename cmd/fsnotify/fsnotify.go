package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				fmt.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("modified:", event.Name)
					if isFileStable(event.Name) {
						fmt.Printf("File %s is stable after modification.\n", event.Name)
						// handleFileStable(event.Name)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("F:\\tmp\\dr")
	if err != nil {
		log.Fatal(err)
	}

	<-done
}

func isFileStable(filePath string) bool {
	const stableDuration = 5 * time.Second
	var lastModTime time.Time
	var lastSize int64
	for {
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			log.Println("error:", err)
			return false
		}
		if !fileInfo.ModTime().Equal(lastModTime) || fileInfo.Size() != lastSize {
			lastModTime = fileInfo.ModTime()
			lastSize = fileInfo.Size()
			// time.Sleep(stableDuration)
		} else {
			return true
		}
		// lastModTime = fileInfo.ModTime()
		// lastSize = fileInfo.Size()
		// // createdTime = fileInfo.CreatedTime()
		// log.Println("ModTime", lastModTime)
		// // log.Println("CreatedTime:", createdTime)
		// log.Println("lastSize:", lastSize)
	}
}

func handleFileStable(filePath string) {
	fmt.Printf("File %s has been successfully transferred.\n", filePath)
	// 这里可以添加你想要执行的代码，例如处理传输完成的文件
}
