package models

import(
  "net/http"
)

type Application interface {
  // Name of the server instance
  Title() string

  // Description of the server instance
  Description() string

  // List of view templates
  Templates(string, string) []*Typeroute
}

// Function signature of the callbacks from the router
type ControllerFunc func(w http.ResponseWriter, r *http.Request)
