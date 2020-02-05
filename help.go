package hanabi

import (
	"fmt"
	"log"

	"github.com/miRemid/amy/tserver/event"
	"github.com/miRemid/hanabi/config"
)

// Help is the help plugin
type Help struct {
	Cmd     string `hana:"help" role:"7"`
	Plugins map[string]p
}

// Parse 解析函数
func (h Help) Parse(evt event.CQSession) {
	p := h.Plugins
	msg := config.NAME + "\n"
	for name, plugin := range p {
		msg += fmt.Sprintf("%s:\t%s\n", name, plugin.plugin.Help())
	}
	log.Println(msg)
	res, err := evt.Send(msg[:len(msg)-1], true, false)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(res.ID)
	}
}

// Help return the plugin's use introduction
func (h Help) Help() string {
	return "查看所有命令使用方式"
}
