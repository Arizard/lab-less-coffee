package services

type ListenerErrorCode int

const (
	ListenerErrorNone ListenerErrorCode = iota
	ListenerErrorInternal
	ListenerErrorBadRequest
	ListenerErrorBadAuth
	ListenerErrorNotFound
)

type Response struct {
	Error     error
	ErrorCode ListenerErrorCode
	Data      interface{}
}

// Listener is an interface which translates between infrastructure (http) and
// application data.
type Listener interface {
	Listen()
}

