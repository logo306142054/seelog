package seelog

import (
	"fmt"
	"log"
	"os"

	"github.com/hpcloud/tail"
)

// 监控日志文件
func monitor(filePath string) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("[seelog] monitor error:%+v", err)
		}
	}()

	t, _ := tail.TailFile(filePath, tail.Config{Follow: true, ReOpen: true, Location: &tail.SeekInfo{
		Offset: 0,
		Whence: os.SEEK_END,
	}})

	for line := range t.Lines {
		fmt.Println(line.Text)
		manager.broadcast <- []byte(line.Text)
	}
}
