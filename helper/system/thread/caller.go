package thread

// 线程工具类

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func GetCaller(skip int) string {
	fn,file, line := FindCaller(skip)
	return fmt.Sprintf("%s:%s:%d", fn,file, line)
}

func getCaller(skip int) (string,string, int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "","", 0
	}
	fcName := runtime.FuncForPC(pc).Name() //获取函数名
	n := 0
	for i := len(file) - 1; i > 0; i-- {
		if string(file[i]) == "/" {
			n++
			if n >= 2 {
				file = file[i+1:]
				break
			}
		}
	}
	return fcName,file, line
}

func FindCaller(skip int) (string,string, int) {
	file := ""
	line := 0
	fn := ""
	for i := 0; i < 10; i++ {
		fn,file, line = getCaller(skip + i)
		if !strings.HasPrefix(file, "log") {
			break
		}
	}
	return fn,file, line
}

//获取go routine id  相当于thread id
func GetRoutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
