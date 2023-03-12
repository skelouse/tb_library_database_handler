package lib

import (
	httpEvent "github.com/taubyte/go-sdk/http/event"
)

func _set(h httpEvent.Event) error {
	_, err := h.Write([]byte("do set"))
	return err
}
