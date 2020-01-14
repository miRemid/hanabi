# Hanabi
一个go语言CQHTTP机器人框架(目前仅支持消息)

# 简单使用
```golang
package main

import (
	"log"

	"github.com/miRemid/hanabi"
	"github.com/miRemid/hanabi/plugins"
)

func main() {
    // 设置Amy SDK参数，CQHTTP地址和端口
    client := hanabi.NewServer("127.0.0.1", 5700)
    // 设置CQHTTP API TOKEN
    client.AccessToken("token")
    
    // 注册插件
    client.Register(plugins.Roll{
        Cmd: "roll",
        Area: 100,
    })

    // hanabi自带了一个Help插件
    // 如需关闭请设置HelpPlugin为false
    // client.HelpPlugin = false

    log.Println("run at 3000")
    // 运行
    client.Run(":3000", "/")
}
```
# 配置
    name: Bot名称=>"amy"
    log_path: 日志存储目录=>"./log"
    screct: cqhttp配置=>"amy"
    access_token: cqhttp api token=>"asdf"
    cmd: 命令标志=>["!", "#"]
# 插件
## Plugin 插件接口
hanabi的插件需要满足`hanabi.Plugin`接口
```golang

import "github.com/miRemid/amy/tserver/event"

type Plugin interface {
    // 作为解析命令函数
    Parse(evt event.CQSession)
    // 返回插件使用方式信息
    Help() string
}
```
## Plugin 插件标准
hanabi的插件标准格式非常简单，其必须包含`Cmd`字段，返回值为`string`用于标识指令.
如果在`Cmd`字段添加`tag`如`hana:"cmd"`，在`Register`时就不必要写入Cmd信息.
```golang
type Example struct {
    Cmd string `hana:"haha" role:"7"`
}

func (e Example) Parse(evt event.CQSession) {
    if res, err := evt.Send("hahaha", true, false); err != nil {
        ...
    }else {
        ...
    }
}

func (e Example) Help() string {
    return "return hahaha"
}
```
## Plugin 响应消息权限
默认所有命令响应群组、私人、讨论组消息。如需更改权限，在插件的Cmd字段中添加tag:`role`.
```golang
type Example struct {
    Cmd string `hana:"example" role:"7"`
}
```
hanabi将会提取role字段转为int类型取低三位数字，根据其二进制判断消息权限

    000 不响应任何消息
    001 私人消息，1
    010 群组消息，2
    100 讨论祖消息，4
    111 所有消息，7