package lib

import (
	"io"

	"github.com/taubyte/go-sdk/database"
	httpEvent "github.com/taubyte/go-sdk/http/event"
)

//go:generate go get github.com/mailru/easyjson
//go:generate go install github.com/mailru/easyjson/...@latest
//go:generate easyjson -all ${GOFILE}

// Used only for un-marshalling the event body onto
type BodyValue struct {
	Value []byte
}

func open() (database.Database, error) {
	return database.New(databaseMatch)
}

func queryKey(h httpEvent.Event) (string, error) {
	return h.Query().Get("key")
}

func bodyValue(h httpEvent.Event) ([]byte, error) {
	defer h.Body().Close()

	data, err := io.ReadAll(h.Body())
	if err != nil {
		return nil, err
	}

	var value BodyValue
	err = value.UnmarshalJSON(data)
	if err != nil {
		return nil, err
	}

	return value.Value, nil
}
