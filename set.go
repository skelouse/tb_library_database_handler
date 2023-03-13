package lib

import (
	"fmt"

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

	err = db.Put(key, value)
	if err != nil {
		return err
	}

	_, err = h.Write([]byte(fmt.Sprintf("put into key `%s`: %s", key, string(value))))
	return err
}
