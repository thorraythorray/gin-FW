// package main

// import (
// 	"flag"
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	"github.com/fsnotify/fsnotify"
// )

// var (
// 	watchPath = flag.String("path", "/data/files/dr/tmp/", "watch file path")
// )

// func main() {
// 	watcher, err := fsnotify.NewWatcher()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer watcher.Close()

// 	done := make(chan bool)
// 	go func() {
// 		for {
// 			select {
// 			case event, ok := <-watcher.Events:
// 				if !ok {
// 					return
// 				}
// 				if event.Op&fsnotify.Create == fsnotify.Create {
// 					// 文件创建事件，可能是传输完成的标志
// 					log.Println("File created:", event.Name)
// 					// 检查文件大小是否稳定，以判断传输是否完成
// 					checkFileTransferComplete(event.Name)
// 				}
// 			case err, ok := <-watcher.Errors:
// 				if !ok {
// 					return
// 				}
// 				log.Println("error:", err)
// 			}
// 		}
// 	}()

// 	err = watcher.Add(*watchPath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	<-done
// }

// func checkFileTransferComplete(filePath string) {
// 	// 检查文件大小是否稳定，以判断传输是否完成
// 	for {
// 		time.Sleep(1 * time.Second) // 等待1秒以确保文件大小稳定
// 		fileInfo, err := os.Stat(filePath)
// 		if err != nil {
// 			log.Println("Error getting file info:", err)
// 			return
// 		}
// 		if fileInfo.Size() == 0 {
// 			log.Println("File size is zero, waiting for transfer to complete...")
// 			continue
// 		}
// 		// 检查文件大小是否在两次检查之间没有变化
// 		prevSize := fileInfo.Size()
// 		time.Sleep(1 * time.Second)

// 		newFileInfo, err := os.Stat(filePath)
// 		if err != nil {
// 			log.Println("Error getting file info:", err)
// 			return
// 		}
// 		if newFileInfo.Size() == prevSize {
// 			log.Printf("File transfer complete: %s\n", filePath)
// 			break
// 		}
// 	}
// }

// func handleFileStable(filePath string) {
// 	fmt.Printf("File %s has been successfully transferred.\n", filePath)
// 	// 这里可以添加你想要执行的代码，例如处理传输完成的文件
// }

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
)

var (
	watchPath = flag.String("path", "E:\\Tmp\\dr", "watch file path")
)

// 检查文件是否“完成”传输的简单逻辑（基于文件大小在一段时间内是否变化）
func isFileTransferred(filePath string, lastSize int64, checkInterval time.Duration) bool {
	fi, err := os.Stat(filePath)
	if err != nil {
		log.Printf("Error getting file info for %s: %v", filePath, err)
		return false
	}
	if fi.Size() != lastSize {
		// 文件大小改变，认为文件还在传输中
		return false
	}
	// 等待一段时间再次检查
	time.Sleep(checkInterval)
	// 再次检查文件大小
	fi2, err := os.Stat(filePath)
	if err != nil {
		log.Printf("Error getting file info for %s: %v", filePath, err)
		return false
	}
	// 如果文件大小在指定时间内没有变化，则认为文件传输完成
	fmt.Println(fi2.Size(), lastSize)
	return fi2.Size() == lastSize
}

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
				if event.Op&fsnotify.Write == fsnotify.Write {
					// 文件被写入，可能是传输正在进行
					filePath := event.Name
					fi, err := os.Stat(filePath)
					if err != nil {
						log.Printf("Error getting file info for %s: %v", filePath, err)
						continue
					}
					if isFileTransferred(filePath, fi.Size(), 500*time.Microsecond) {
						// 文件传输完成，执行其他操作
						fmt.Printf("File %s has been transferred.\n", filePath)
						// TODO: 在这里添加你的操作
						// err := os.Remove(filePath)
						// if err != nil {
						// 	fmt.Printf("Error deleting file %s: %v\n", filePath, err)
						// 	continue
						// }
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

	// 添加要监控的目录
	err = watcher.Add(*watchPath)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

// 注意：这个示例中的isFileTransferred函数是简化的，并且依赖于特定的逻辑（文件大小在一段时间内不变）。
// 在实际应用中，你可能需要更复杂的逻辑来准确判断文件是否传输完成。
