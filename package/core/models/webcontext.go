package coremodels

import(
  "net/http"
)

type Webcontext struct {
  W http.ResponseWriter
  R *http.Request
}

func NewWebcontext(w http.ResponseWriter, r *http.Request) *Webcontext {
  return &Webcontext{
    W: w,
    R: r,
  }
}
