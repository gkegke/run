/*

  Helper functions and structs related to handlers.
  
   Error

   StatusError
     - Error()
     - Status()

   Env

   Handler
     - ServeHTTP

*/

package main

import (
    "log"

    "net/http"
)

type Error interface {
    error
    Status() int
}

type StatusError struct {
    Code int
    Err error
}

func (se StatusError) Error() string {
    return se.Err.Error()
}

func (se StatusError) Status() int {
    return se.Code
}

type Env struct {
    Port string
}

/*

  An easier too use handler that makes handling errors and
  returning http statuses easier.

  Main use case is it removes the need too return, when an error
  occurs.

*/
type Handler struct {
    *Env
    H func(e *Env, w http.ResponseWriter, r *http.Request) error
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    err := h.H(h.Env, w, r)
    if err != nil {
        switch e := err.(type) {
        case Error:
            log.Printf("HTTP %d - %s", e.Status(), e)
            http.Error(w, e.Error(), e.Status())
        default:
            http.Error(w, http.StatusText(http.StatusInternalServerError),
                          http.StatusInternalServerError)
        }
    }
}
