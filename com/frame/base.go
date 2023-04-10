package frame

import (
	"net/http"
)

type BaseAction interface {
	Invoke() error
	Register() (http.Handler, error)
}
