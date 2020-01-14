package hanabi

import "testing"

import "github.com/miRemid/amy/tserver/event"

type echo struct {
	Cmd string `hana:"echo" role:"7"`
}

func (e echo) Parse(evt event.CQSession) {
	evt.Send(evt.RawMessage, true, true)
}

func (e echo) Help() string {
	return "Echo"
}

func TestRun(t *testing.T) {
	client := NewServer("127.0.0.1", 5700)
	client.Register(echo{
		Cmd: "echo",
	})
	client.HelpPlugin = false
	client.Run(":8888", "/")
}
