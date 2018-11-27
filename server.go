package seelog

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"runtime"
	"time"

	"golang.org/x/net/websocket"
)

// 开启 httpServer
func server(port int) {

	defer func() {
		if err := recover(); err != nil {
			log.Printf("[seelog] error:%+v", err)
		}
	}()

	// 返回页面
	http.HandleFunc("/", page)
	// socket链接
	http.Handle("/ws", websocket.Handler(genConn))
	// 测试
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		_, currentfile, _, _ := runtime.Caller(0)
		filename := path.Join(path.Dir(currentfile), "index.html")
		t, err := template.ParseFiles(filename)
		if err != nil {
			log.Println(err)
		}
		t.Execute(writer, nil)
	})
	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

// 创建client对象
func genConn(ws *websocket.Conn) {
	client := &client{time.Now().String(), ws, make(chan []byte, 1024)}
	manager.register <- client
	go client.read()
	client.write()
}
