package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {

	filepath := "log/"

	logFile, err := os.OpenFile(filepath+"log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Print(err)
	}
	logger := logrus.New()
	// 将日志输出
	logger.Out = logFile
	logger.SetLevel(logrus.DebugLevel)
	// 分割日志
	logWriter, _ := retalog.New(
		// 文件目录
		filepath+"%Y-%-m%d.log",
		// 删除时间一周
		retalog.WithMaxAge(7*24*time.Hour),
		// 分割时间 一天
		retalog.WithRotationTime(24*time.Hour),
	)
	writerMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writerMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// 添加钩子函数
	logger.AddHook(Hook)

	return func(c *gin.Context) {
		// 记录访问开始的时间
		startTime := time.Now()
		// 执行后续拦截器
		c.Next()
		endTime := time.Since(startTime).Milliseconds()
		// 花费的时间,毫秒
		spendTime := fmt.Sprintf("%d ms", endTime)
		// 获取主机名
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknow"
		}
		// 请求状态码
		status := c.Writer.Status()
		// 来源ip
		clientIp := c.ClientIP()
		// userAgent
		userAgent := c.Request.UserAgent()
		// 数据大小
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		// 请求方法
		method := c.Request.Method
		// 路由路径
		path := c.Request.RequestURI
		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"status":    status,
			"SpendTime": spendTime,
			"Ip":        clientIp,
			"Method":    method,
			"Path":      path,
			"DataSize":  dataSize,
			"Agent":     userAgent,
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())

		}
		if status >= 500 {
			entry.Error()
		} else if status >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}

	}
}
