package action

import (
	"log"
	"net/http"
)

type GetBookList struct {
}

func (self *GetBookList) Invoke() (string, error) {
	log.Printf("/getbooklist")
	return "hello api[GetBookList]", nil
}

func (self *GetBookList) Register() (http.Handler, error) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		out, err := self.Invoke()
		if err != nil {
			log.Fatal(err)
		}

		w.Write([]byte(out))
	}

	return http.HandlerFunc(handler), nil
}
