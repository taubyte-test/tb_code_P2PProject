package lib

import (
	"fmt"
    "bitbucket.org/taubyte/go-sdk/event"
	"bitbucket.org/taubyte/go-sdk/p2p/node"
)

//export ping
func ping(e event.Event) uint32 {
    h, err := e.HTTP()
	if err != nil {
		return 1
	}

	err = runPing(h)
	if err != nil {
		errString := fmt.Sprintf(`{"error": ping failed with %s}`, err)
		h.Write([]byte(errString))
		return 1
	}

	return 0
}

func runPing(h event.HttpEvent) error {
	cmd, err := node.New("/test/v1").Command("ping")
	if err != nil {
		return err
	}

	response, err := cmd.Send([]byte("Hello, world"))
	if err != nil {
		return err
	}

	_, err = h.Write(response)
	return err
}