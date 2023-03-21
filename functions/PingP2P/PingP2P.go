package lib

import (
	"fmt"

	"github.com/taubyte/go-sdk/event"
	p2pEvent "github.com/taubyte/go-sdk/p2p/event"
)

//export ping
func ping(e event.Event) uint32 {
	p, err := e.P2P()
	if err != nil {
		return 1
	}

	err = runPing(p)
	if err != nil {
		errString := fmt.Sprintf(`{"error": "ping failed with %s"}`, err)
		p.Write([]byte(errString))
		return 1
	}

	return 0
}

func runPing(e p2pEvent.Event) error {
	command, err := e.Command()
	if err != nil {
		return err
	}

	data, err := e.Data()
	if err != nil {
		return err
	}

	from, err := e.From()
	if err != nil {
		return err
	}

	protocol, err := e.Protocol()
	if err != nil {
		return err
	}

	to, err := e.To()
	if err != nil {
		return err
	}

	toWrite := fmt.Sprintf(`{
		"protocol": "%s",
		"command": "%s",
		"data": "%s",
		"from": "%s",
		"to": "%s"
}`, protocol, command, string(data), from.String(), to.String())

	return e.Write([]byte(toWrite))
}
