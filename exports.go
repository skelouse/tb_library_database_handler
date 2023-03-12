package lib

import (
	"fmt"

	httpEvent "github.com/taubyte/go-sdk/http/event"
)

func handleError(h httpEvent.Event, err error) uint32 {
	if err != nil {
		h.Return(418)

		_, err := h.Write([]byte(fmt.Sprintf("failed with: %s", err.Error())))
		if err != nil {
			h.Return(500)
		}
	}

	return 0
}

//export get
func get(h httpEvent.Event) uint32 {
	return handleError(h, _get(h))
}

//export set
func set(h httpEvent.Event) uint32 {
	return handleError(h, _set(h))
}

//export delete
func delete(h httpEvent.Event) uint32 {
	return handleError(h, _delete(h))
}

//export list
func list(h httpEvent.Event) uint32 {
	return handleError(h, _list(h))
}
