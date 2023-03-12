package lib

import (
	httpEvent "github.com/taubyte/go-sdk/http/event"
)

func _get(h httpEvent.Event) error {
	_, err := h.Write([]byte("do get"))
	return err
}
