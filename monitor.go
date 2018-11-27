package seelog

import (
	"log"
	"os"
	"time"

	"github.com/hpcloud/tail"
)

// 监控日志文件
func monitor(filePath string) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("[seelog] error:%+v", err)
		}
	}()

	//fileInfo := checkFile(filePath)
	checkFile(filePath)

	t, _ := tail.TailFile(filePath, tail.Config{Follow: true, ReOpen: true})
	for line := range t.Lines {
		//fmt.Println(line.Text)
		manager.broadcast <- []byte(line.Text)
	}
}

//等待文件被创建
func checkFile(filePath string) os.FileInfo {
	var fileInfo os.FileInfo
	var err error
	for {
		fileInfo, err = os.Stat(filePath)
		if err != nil {
			log.Printf("[seelog] error:%v", err.Error())
			time.Sleep(10 * time.Millisecond)
		} else {
			break
		}
	}
	return fileInfo
}
