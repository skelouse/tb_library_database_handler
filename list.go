package lib

import (
	httpEvent "github.com/taubyte/go-sdk/http/event"
)

func _list(h httpEvent.Event) error {
	_, err := h.Write([]byte("do list"))
	return err
}
