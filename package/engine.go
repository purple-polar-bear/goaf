package apif

import(
  "net/http"
  "strings"
  "oaf-server/package/controllers"
)

// A controller for resolving the OGC Api Feature calls
//
// The controller contains all the elements required
// for handling OGC Api Features calls.
type Engine interface {
  HTTPHandler(http.ResponseWriter, *http.Request)
  Mount(mountingPath string)
}

type engine struct {
  // Mounting path is the path where the controller is mounted.
  //
  // Example:
  mountingPath string

  router Router
}

// Function signature of the callbacks from the router
type ControllerFunc func(w http.ResponseWriter, r *http.Request)

// Returns a new controller for handling OGC Api Feature calls.
//
// This method does not require any furhter configuration
func NewEngine(router Router) *engine {
  engine := &engine{
    router: router,
  }
  return engine
}

func NewSimpleEngine(mountingPath string) *engine {
  router := NewRouter(mountingPath)
  engine := NewEngine(router)

  controller := apifcontrollers.LandingPage{}
  router.AddRoute("", controller.Handle)
  router.AddRoute("/", controller.Handle)
  router.AddRoute("/conformance", controller.Handle)
  router.AddRoute("/collections", controller.Handle)

  return engine
}

func (c *engine) HTTPHandler(w http.ResponseWriter, r *http.Request) {
  if (c.router == nil) {
    panic("Apif controller is not mounted")
  }

  c.router.HandleRequest(w, r)
}

func (c *engine) Mount(mountingPath string) {
  mountingPath = sanitizeMountingPath(mountingPath)
  c.mountingPath = mountingPath
}

func sanitizeMountingPath(mountingPath string) string {
  if(!strings.HasPrefix(mountingPath, "/")) {
    mountingPath = "/" + mountingPath
  }

  return mountingPath
}
