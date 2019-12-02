package plugins

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"

	"github.com/miRemid/amy/tserver/event"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// Roll 随机数插件
type Roll struct {
	Cmd  string `hana:"cmd" role:"7"`
	Area int
}

func (r Roll) Parse(evt event.CQSession) {
	if r.Area == 0 {
		r.Area = 100
	}
	text := fmt.Sprintf("%d", rand.Intn(r.Area))
	evt.Send(text, true, true)
}

func (r Roll) Help() string {
	v := reflect.ValueOf(&r.Area)
	if r.Area == 0 {
		v.Elem().SetInt(100)
	}
	return fmt.Sprintf("随机生成一个数字(0-%d)", r.Area)
}
