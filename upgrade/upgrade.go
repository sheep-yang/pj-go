package upgrade

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

//WhiteList 定义白名单用户信息
var WhiteList = map[int]struct{}{}

func main() {
	//WhiteFunc()
	IsPass(1000)

}

//checkErr 错误处理
func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

// init 获取白名单用户信息,生成用户信息;
func init() {
	path := "you path"
	fileinfo, err := os.Open(path)
	checkErr(err)
	defer fileinfo.Close()
	rd := bufio.NewReader(fileinfo)
	for {
		line, _, err := rd.ReadLine()
		if err != nil || io.EOF == err {
			break
		}
		userid, err := strconv.Atoi(string(line))
		checkErr(err)
		WhiteList[userid] = struct{}{}
	}

}

//IsPass 判断用户是否在白名单
func IsPass(userid int) bool {
	if _, ok := WhiteList[userid]; ok {
		return true
	}
	return false
}
