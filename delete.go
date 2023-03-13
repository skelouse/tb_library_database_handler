package lib

import (
	"github.com/taubyte/go-sdk/database"
	httpEvent "github.com/taubyte/go-sdk/http/event"
)

func _delete(h httpEvent.Event) error {
	db, err := database.New(databaseMatch)
	if err != nil {
		return err
	}

	key, err := h.Query().Get("key")
	if err != nil {
		return err
	}

	return db.Delete(key)
}
