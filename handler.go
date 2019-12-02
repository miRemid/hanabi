package hanabi

import (
	"log"

	"github.com/miRemid/amy/tserver/event"
	"github.com/miRemid/hanabi/config"
)

func (c *Client) m(evt event.CQSession) {
	cmd, _ := evt.Params(config.CMD...)
	log.Println(cmd)
	if plugin, ok := c.plugins[cmd]; ok {
		flag := false
		switch evt.Type {
		case "private":
			if plugin.private {
				log.Printf("[I] 私人消息，响应%s插件", plugin.name)
				flag = true
			}
			break
		case "group":
			if plugin.group {
				log.Printf("[I] 群组消息，响应%s插件", plugin.name)
				flag = true
			}
			break
		case "discuss":
			if plugin.discuss {
				log.Printf("[I] 讨论族消息，响应%s插件", plugin.name)
				flag = true
			}
			break
		}
		if flag {
			plugin.plugin.Parse(evt)
		}
	}
}
