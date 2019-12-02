package hanabi

import (
	"fmt"
	"log"
	"reflect"
	"strconv"

	"github.com/miRemid/amy/tserver"
	"github.com/miRemid/amy/tserver/event"

	"github.com/miRemid/hanabi/config"
)

const (
	// Notice is the flag of notice Parse function
	Notice = iota
	// Request is the flag of request Parse function
	Request
)

type Plugin interface {
	Parse(evt event.CQSession)
	Help() string
}

type Handler func(evt event.CQSession)

// Parse 解析函数
type Parse func(evt event.CQSession)

// Client is the main struct
type Client struct {
	bot     *tserver.Bot
	plugins map[string]p
	notice  Handler
	request Handler
}

type p struct {
	plugin  Plugin
	name    string
	private bool
	group   bool
	discuss bool
}

// NewServer return a hanabi client
func NewServer(url string, port	int) *Client {
	var res Client
	res.bot = tserver.NewBot(url, port)
	res.plugins = make(map[string]p)
	return &res
}

func permision(v reflect.Value, f reflect.StructField) int {
	role := f.Tag.Get("role")
	if role == "" {
		log.Printf("[WA] %s插件权限读取失败，以设置默认权限为7", v.Type())
		return 7
	}
	if per, err := strconv.Atoi(role); err != nil {
		log.Printf("[WA] %s插件权限读取失败，以设置默认权限为7", v.Type())
		return 7
	} else {
		return per
	}
}

// Register a plugin
func (c *Client) Register(pluginss ...Plugin) {
	for _, plugin := range pluginss {
		var cmd string
		var role int
		v := reflect.ValueOf(plugin)
		t := reflect.TypeOf(plugin)
		if f, ok := t.FieldByName("Cmd"); !ok {
			log.Printf("[E] %s插件读取失败，检查是否包含Cmd字段", v.Type())
			continue
		} else {
			cmd = fmt.Sprintf("%s", v.FieldByName("Cmd"))
			if cmd == "" {
				cmd = f.Tag.Get("hana")
			}
			role = permision(v, f)
		}
		if cmd == "" {
			log.Printf("[E] %s插件读取失败，检查初始化是否正确或tag是否包含hana字段", v.Type())
			continue
		}
		if _, ok := c.plugins[cmd]; ok {
			log.Printf("[P] %s插件覆盖成功", v.Type())
		} else {
			log.Printf("[P] %s插件读取成功", v.Type())
		}
		private, group, discuss := getp(role)
		c.plugins[cmd] = p{
			plugin:  plugin,
			name:    fmt.Sprintf("%s", v.Type()),
			private: private,
			group:   group,
			discuss: discuss,
		}
	}
}

func getp(role int) (bool, bool, bool) {
	private := role & 1
	group := role >> 1 & 1
	discuss := role >> 2 & 1
	return private == 1, group == 1, discuss == 1
}

// On set the handler function
func (c *Client) On(handler Handler, flag int) {
	switch flag {
	case Notice:
		c.notice = handler
		break
	case Request:
		c.request = handler
		break
	default:
		break
	}
}

// AccessToken set the CQHTTP api token
func (c *Client) AccessToken(token string) {
	c.bot.AccessToken = token
}

// Run the client
func (c *Client) Run(addr, router string) {
	c.Register(Help{
		Cmd:     "help",
		Plugins: c.plugins,
	})
	if config.SCRECT != "" {
		c.bot.Use(tserver.Signature(config.SCRECT))
	}
	c.bot.On(c.m, tserver.Message)
	c.bot.Run(addr, router)
}
