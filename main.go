package main

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gom3u8/conf"
	_ "gom3u8/data"
	"gom3u8/task"
	"gom3u8/work"
	"gopkg.in/natefinch/lumberjack.v2"
	"io/fs"
	"net/http"
	"strconv"
)

//go:embed static
var static embed.FS

func main() {
	conf.ConfInit()
	logFile := &lumberjack.Logger{
		Filename:   "./log/log.txt",
		MaxSize:    10, // MB
		MaxBackups: conf.ConfMap.Log.LogNu,
		MaxAge:     28, // days
		Compress:   true,
		LocalTime:  true,
	}

	defer logFile.Close()

	log.SetOutput(logFile)

	run()

}
func run() {
	go work.Working()
	r := gin.Default()
	tc := task.TaskController{}
	sub, err := fs.Sub(static, "static")
	if err != nil {
		panic(err)
	}
	r.StaticFS("/static", http.FS(sub))
	r.POST("/addTask", tc.AddTask)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})
	fmt.Println("open http://127.0.0.1:" + strconv.Itoa(conf.ConfMap.Init.Port) + "/static/ ")
	r.Run(":" + strconv.Itoa(conf.ConfMap.Init.Port)) // 监听2045端口
}
