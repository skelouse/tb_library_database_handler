package lib

import (
	"strings"

	"github.com/taubyte/go-sdk/database"
	httpEvent "github.com/taubyte/go-sdk/http/event"
)

func _list(h httpEvent.Event) error {
	db, err := database.New(databaseMatch)
	if err != nil {
		return err
	}

	key, err := h.Query().Get("key")
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
