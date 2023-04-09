package frame

import (
	"net/http"
)

type BaseAction interface {
	inv() error
	Register() (http.Handler, error)
}
