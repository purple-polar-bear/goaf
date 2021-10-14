package apif

import(
  "net/http"
  "strings"
  "oaf-server/package/controllers"
  "oaf-server/package/models"
)

// A controller for resolving the OGC Api Feature calls
//
// The controller contains all the elements required
// for handling OGC Api Features calls.
type Engine interface {
  // HTTP functions

  HTTPHandler(http.ResponseWriter, *http.Request)
  Mount(mountingPath string)

  // Template functions

  AddConformanceTemplate(title, contentType, WFSConformanceTemplater)

  // Service functions
  Title() string
  Description() string
  SetTitle(title string)
  SetDescription(description string)
  Templates() []models.Template
}

type engine struct {
  // Mounting path is the path where the controller is mounted.
  //
  // Example:
  mountingPath string

  // router
  router Router

  // landingpage controller
  landingpageController apifcontrollers.LandingpageController

  //conformance controller
  conformanceController apifcontrollers.ConformanceController

  // conformance templates
  conformanceTemplates map[string]templates.RenderConformanceFunc

  title string
  description string
}

// Function signature of the callbacks from the router
type ControllerFunc func(w http.ResponseWriter, r *http.Request)

// Returns a new controller for handling OGC Api Feature calls.
//
// This method does not require any furhter configuration
func NewEngine(router Router) *engine {
  engine := &engine{
    router: router,
    conformanceTemplates: make*map[string]templates.RenderConformanceFunc,
  }
  return engine
}

func NewSimpleEngine(mountingPath string) *engine {
  router := NewRouter(mountingPath)
  engine := NewEngine(router)

  engine.landingpageController := apifcontrollers.LandingPage{}
  router.AddRoute("", "/?", landingpageController.Handle)

  engine.conformanceController := apifcontrollers.ConformanceController{}
  router.AddRoute("conformance", "/conformance", conformanceController.Handle)


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

func (e *engine) Title() { return e.title }
func (e *engine) Description() { return e.description }
func (e *engine) SetTitle(title string) { e.title = title }
func (e *engine) SetDescription(description string) { e.description = description }

func (e *engine) Templates(url string, contenttype string) []models.Template {

}

//
// Utility functions
//

func sanitizeMountingPath(mountingPath string) string {
  if(!strings.HasPrefix(mountingPath, "/")) {
    mountingPath = "/" + mountingPath
  }

  return mountingPath
}
