package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"syscall"
)

const LocalDateTimeFormat string = "2006-01-02 15:04:05"

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func IF(test bool, then, other interface{}) interface{} {
	if test {
		return then
	}
	return other
}

// TryCatch 实现函数的try...catch实现
func TryCatch(fun func(), handler func(i interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

// StructToMap 结构体转换为Map
func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[strings.ToLower(t.Field(i).Name)] = v.Field(i).Interface()
	}
	return data
}

//注册退出信号
func HandleExitSignal(exitFunc func()) {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	for {
		switch <-signalChan {
		case syscall.SIGINT:
			fallthrough
		case syscall.SIGKILL:
			fallthrough
		case syscall.SIGTERM:
			if nil != exitFunc {
				TryCatch(exitFunc, func(i interface{}) {
					os.Exit(0)
				})
			}
			log.Println(os.Args[0], "exited!")
			os.Exit(0)
		}
	}
}

func HttpRequest(method, url, params string, header map[string]string) (map[string]interface{}, error) {
	http_client := &http.Client{}
	body := strings.NewReader(params)

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	resp, err := http_client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	r, err := ioutil.ReadAll(resp.Body)
	result := make(map[string]interface{})
	if err := json.Unmarshal(r, &result); err != nil {
		return nil, err
	}
	return result, nil
}
