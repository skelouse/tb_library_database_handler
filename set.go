package lib

import (
	"fmt"
	"io"

	"github.com/taubyte/go-sdk/database"
	httpEvent "github.com/taubyte/go-sdk/http/event"
)

func _set(h httpEvent.Event) error {
	db, err := database.New(databaseMatch)
	if err != nil {
		return err
	}

	key, err := h.Query().Get("key")
	if err != nil {
		return err
	}

	defer h.Body().Close()
	data, err := io.ReadAll(h.Body())
	if err != nil {
		return err
	}

	fmt.Println("DATA", string(data))

	var value BodyValue
	err = value.UnmarshalJSON(data)
	if err != nil {
		return err
	}

	err = db.Put(key, []byte(value.Value))
	if err != nil {
		return err
	}

	_, err = h.Write([]byte(fmt.Sprintf("put into key `%s`: %s", key, value.Value)))
	return err
}
