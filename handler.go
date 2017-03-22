package honey

import (
	"log"
	"net/http"
	"time"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func Handle(f HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var status int
		start := time.Now()

		defer func() {
			if e := recover(); e != nil {
				status = 500
			}

			log.Printf("\"%s %s\" %d %d \"%s\"", r.Method, r.RequestURI, status, time.Since(start).Nanoseconds()/1000000, r.UserAgent())
		}()

		err := f(w, r)
		if err != nil {
			log.Printf("[ERROR] %v", err)
			status = http.StatusInternalServerError
		} else {
			status = http.StatusOK
		}
	}
}
