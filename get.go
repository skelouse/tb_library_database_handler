package lib

import (
	httpEvent "github.com/taubyte/go-sdk/http/event"
)

func _get(h httpEvent.Event) error {
	db, err := open()
	if err != nil {
		return err
	}

	key, err := queryKey(h)
	if err != nil {
		return err
	}

	value, err := db.Get(key)
	if err != nil {
		return err
	}

	_, err = h.Write(value)
	return err
}
