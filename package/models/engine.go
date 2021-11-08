package models

import(
  "net/http"
)

type Application interface {
  // Configuration of the server
  Config() Serverconfig

  // List of routes
  Routes() []Route

  // List of view templates
  Templates(string, string) []Handler

  GetService(string) interface{}
}

// Function signature of the callbacks from the router
type ControllerFunc func(w http.ResponseWriter, r *http.Request)
