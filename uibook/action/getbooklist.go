package action

import (
	"log"
	"net/http"
)

type GetBookList struct {
}

func (self *GetBookList) inv() error {
	log.Printf("hello api[GetBookList]")
	return nil
}

func (self *GetBookList) Register() (http.Handler, error) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		self.inv()
	}

	return http.HandlerFunc(handler), nil
}
