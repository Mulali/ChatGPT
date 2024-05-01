package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    // 引入OpenAI包或其他相关依赖
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

// 假设你有一个函数用来和OpenAI进行交互
func chatWithGPT(input string) (string, error) {
    // 交互逻辑
    return "", nil
}

func handleConnections(c *gin.Context) {
    // 升级get请求为websocket
    ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        c.AbortWithStatus(http.StatusInternalServerError)
        return
    }
    defer ws.Close()

    for {
        // 读取消息
        messageType, p, err := ws.ReadMessage()
        if err != nil {
            c.AbortWithStatus(http.StatusInternalServerError)
            return
        }

        // 处理消息，与GPT对话
        response, err := chatWithGPT(string(p))
        if err != nil {
            c.AbortWithStatus(http.StatusInternalServerError)
            return
        }

        // 发送回复
        if err := ws.WriteMessage(messageType, []byte(response)); err != nil {
            c.AbortWithStatus(http.StatusInternalServerError)
            return
        }
    }
}

func main() {
    router := gin.Default()
    router.GET("/ws", handleConnections)
    router.Run(":8080")
}
