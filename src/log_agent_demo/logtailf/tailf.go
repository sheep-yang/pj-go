package logtailf

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

var (
	tails   *tail.Tail
	Msgchan = make(chan string, 1024) //定义一个通道存储msg
	err     error
	Msg     *tail.Line
	ok      bool
)

func Initlog(filepath string) {
	tails, err = tail.TailFile(filepath, tail.Config{
		ReOpen: true,
		Follow: true,
		//	Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}
	fmt.Println("初始化tailf成功")
}

//从文件读取内容
func Readlog() {
	for {
		Msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println("发送的消息Msg.Text:", Msg.Text)
		Msgchan <- Msg.Text
	}

}
