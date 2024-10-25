package main

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
	"sync"
	"time"
)

var (
	x int64

	wg sync.WaitGroup // 等待组

	m sync.Mutex
)

// add 对全局变量x执行5000次加1操作
func add() {
	for i := 0; i < 5000; i++ {
		m.Lock()
		x = x + 1
		m.Unlock()
	}
	wg.Done()
}

func InitLogs() *logrus.Logger {
	// 1、使用默认的logger
	logger := logrus.StandardLogger()
	// 2、对默认logger进行调整设置（设置日志格式）
	logger.SetFormatter(&nested.Formatter{
		NoColors:        true,
		CallerFirst:     true,
		TimestampFormat: time.DateTime,
		HideKeys:        true,
		CustomCallerFormatter: func(frame *runtime.Frame) string {
			filename := path.Base(frame.File)
			return fmt.Sprintf(" [%-20.20s : %-4d]", filename, frame.Line)
		},
	})
	logger.SetReportCaller(true)
	logger.Out = io.MultiWriter(os.Stdout)
	return logger
}

func main() {
	InitLogs()
	wg.Add(2)
	logrus.Infof("开始处理资源文件...")
	i := 0
	for {
		i++
		logrus.Infof("haha")
		if i > 10 {
			break
		}

	}
	go add()
	go add()

	wg.Wait()
	fmt.Println(x)
}
