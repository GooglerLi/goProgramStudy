package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// 定义一个 Upgrader 配置
// 用于将普通的 HTTP 连接升级为 WebSocket 连接
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 在生产环境中，这里应该验证请求的来源以防止 CSRF 攻击
		// 例如：return r.Header.Get("Origin") == "https://yourdomain.com"
		return true // 允许所有跨域请求（仅用于开发测试）
	},
}

// 处理 WebSocket 连接的核心函数
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// 1. 将 HTTP 连接升级为 WebSocket 连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Upgrade failed:", err)
		return
	}
	defer conn.Close() // 确保函数退出时关闭连接

	// 2. 连接建立成功，开始处理消息
	fmt.Println("Client connected!")

	for {
		// 3. 循环读取客户端发送的消息
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read failed:", err)
			break // 如果读取错误（如客户端断开），则退出循环
		}
		log.Printf("Received: %s", message)

		// 4. 处理消息：这里简单地将消息原样发回给客户端（ Echo 服务器 ）
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Write failed:", err)
			break
		}
		// 你也可以在这里根据消息内容进行不同的业务逻辑处理，
		// 例如广播给其他连接、查询数据库等。
	}
	fmt.Println("Client disconnected.")
}

func StartServer() {
	// 注册路由，指定 WebSocket 处理函数
	http.HandleFunc("/ws", handleWebSocket)

	// 启动 HTTP 服务器
	fmt.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
