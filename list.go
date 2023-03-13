package lib

import (
	"strings"

	httpEvent "github.com/taubyte/go-sdk/http/event"
)

func _list(h httpEvent.Event) error {
	db, err := open()
	if err != nil {
		return err
	}

	key, err := queryKey(h)
	if err != nil {
		return err
	}

	keys, err := db.List(key)
	if err != nil {
		return err
	}

	_, err = h.Write([]byte(strings.Join(keys, ",")))
	return err
}
