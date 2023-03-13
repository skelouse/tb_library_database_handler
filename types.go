package lib

//go:generate go get github.com/mailru/easyjson
//go:generate go install github.com/mailru/easyjson/...@latest
//go:generate easyjson -all ${GOFILE}

// Used only for un-marshalling the event body onto
type BodyValue struct {
	Value string
}
