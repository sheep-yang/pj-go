package conf

//Kafaka 定义结构体
type Kafaka struct {
	Ipaddr string `ini:"ipaddr"`
	Topic  string `ini:"topic"`
}

//Tailf 定义路径
type Tailf struct {
	Path string `ini:"path"`
}

//LogConf 定义配置文件信息
type LogConf struct {
	Kafaka   `ini:"kafaka"`
	Tailf    `ini:"tailf"`
}
