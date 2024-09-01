package log

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/wishicorp/sdk/log"
	"io"
	"os"
	"path/filepath"
	"sync"
)

func NewLogger(path, name string, level hclog.Level, console bool, nsq NsqConfig) hclog.InterceptLogger {
	var writers []io.Writer
	var useNsq bool
	useNsq = false
	// 文件流
	if len(path) != 0 && len(name) != 0 {
		logFile, err := NewFileWriter(filepath.Join(path, name) + ".log")
		if err != nil || logFile == nil {
			return nil
		}
		writers = append(writers, logFile)
	}
	// 日志名
	if len(name) == 0 {
		name = "main-loggeer"
	}
	// 控制台
	if console {
		writers = append(writers, os.Stdout)
	}

	// nsq 队列
	if len(nsq.Profile) != 0 && len(nsq.LogName) != 0 && len(nsq.NSQLookupdServer) != 0 {
		useNsq = true
		index := fmt.Sprintf("%s-%s", nsq.LogName, nsq.Profile)
		nsqWriter, err := log.NewNsqWriter([]string{nsq.NSQLookupdServer}, nsq.LogName, index)
		if nil == err {
			writers = append(writers, nsqWriter)
		}
	}

	opts := &hclog.LoggerOptions{
		Name:            name,
		Level:           level,
		Output:          io.MultiWriter(writers...),
		JSONFormat:      useNsq,
		Mutex:           &sync.RWMutex{},
		IncludeLocation: true,
		Exclude: func(level hclog.Level, msg string, args ...interface{}) bool {
			if len(args) == 0 {
				return false
			}
			for _, arg := range args {
				if arg == "schema" {
					return true
				}
			}
			return false
		},
	}
	return hclog.NewInterceptLogger(opts)
}

var Log hclog.InterceptLogger

func SetLogger(l hclog.InterceptLogger) {
	if l != nil {
		Log = l
	}
}
