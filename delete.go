package lib

import (
	httpEvent "github.com/taubyte/go-sdk/http/event"
)

func _delete(h httpEvent.Event) error {
	_, err := h.Write([]byte("do delete"))
	return err
}
