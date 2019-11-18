// 全局接受系统消息（kill -INT 等消息）
//
package notify

import (
	"os"
	"os/signal"
	"time"
)

var AllOfDone = make(chan int, 0)
var notifySubscriibes = make([]func(), 0)

func InitNotify() error {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c

		for _, sub := range notifySubscriibes {
			sub()
		}
		time.Sleep(1 * time.Second) // 等待结果的输出，避免过早结束进程，导致无法看到订阅函数的输出

		close(AllOfDone)
	}()
	return nil
}

func Subscribe(fn func()) {
	notifySubscriibes = append(notifySubscriibes, fn)
}
