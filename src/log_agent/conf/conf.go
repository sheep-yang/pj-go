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

//Etcdinfo 定义配置文件信息
type Etcdinfo struct {
	Etcdipaddr string `ini:"etcdipaddr"`
}

//Esinfo 定义配置文件信息
type Esinfo struct {
	Esurl string `ini:"esurl"`
}

//LogConf 定义配置文件信息
type LogConf struct {
	Kafaka   `ini:"kafaka"`
	Tailf    `ini:"tailf"`
	Etcdinfo `ini:"etcdinfo"`
	Esinfo   `ini:"esinfo"`
}
