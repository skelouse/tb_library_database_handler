package lib

import (
	httpEvent "github.com/taubyte/go-sdk/http/event"
)

func _set(h httpEvent.Event) error {
	db, err := open()
	if err != nil {
		return err
	}

	key, err := queryKey(h)
	if err != nil {
		return err
	}

	value, err := bodyValue(h)
	if err != nil {
		return err
	}

	return db.Put(key, value)
}
