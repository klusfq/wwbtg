package frame

import (
	"net/http"
)

type BaseAction interface {
	Invoke() (string, error)
	Register() (http.Handler, error)
}
