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
  Controller(name string) apifcontrollers.BaseController

  // AddConformanceTemplate(title string, contentType string, renderer coretemplates.RenderConformanceType)
  AddTemplate(name string, title string, contenttype string, renderer interface{})
  Templates(string, string) []*models.Typeroute

  // Service functions
  Title() string
  Description() string
  SetTitle(title string)
  SetDescription(description string)
}

type engine struct {
  // Mounting path is the path where the controller is mounted.
  //
  // Example:
  mountingPath string

  // router
  router Router

  // list of controllers
  // path: routename -> controller
  // controllers map[string]apifcontrollers.BaseController
  // list render functions
  // path: routename -> content type -> controller handler function with renderer

  title string
  description string
}

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

  landingpageController := &apifcontrollers.LandingpageController{}
  engine.AddRoute("", "/?", landingpageController)

  conformanceController := apifcontrollers.NewConformanceController()
  engine.AddRoute("conformance", "/conformance", conformanceController)

  return engine
}

func EnableFeatures(engine *engine) {
  engine.AddConformanceClass("http://www.opengis.net/spec/ogcapi-features-1/1.0/conf/features")
  collectionsController := &apifcontrollers.CollectionsController{}
  engine.AddRoute("collections", "/collection", collectionsController)
}

/*
func EnableTiles(engine *engine) {
  engine.AddConformanceClass("http://www.opengis.net/spec/ogcapi-features-1/1.0/conf/tiles")
  engine.tilesomethingController = apifcontrollers.TilesomethingController{}
  router.AddRoute("tiles", "/tiles", engine.tilesomethingController)
}
*/

func (c *engine) HTTPHandler(w http.ResponseWriter, r *http.Request) {
  if (c.router == nil) {
    panic("Apif controller is not mounted")
  }

  c.router.HandleRequest(w, r)
}

func (e *engine) AddRoute(routename string, path string, controller apifcontrollers.BaseController) {
  e.router.AddRoute(routename, path, controller)
}

func (c *engine) Mount(mountingPath string) {
  mountingPath = sanitizeMountingPath(mountingPath)
  c.mountingPath = mountingPath
}

func (e *engine) Title() string { return e.title }
func (e *engine) Description() string { return e.description }
func (e *engine) SetTitle(title string) { e.title = title }
func (e *engine) SetDescription(description string) { e.description = description }
func (e *engine) Templates(url string, contenttype string) []*models.Typeroute {
  return []*models.Typeroute{}
}

func (e *engine) Controller(name string) apifcontrollers.BaseController {
  return e.router.Controller(name)
}

func (e *engine) AddConformanceClass(conformanceclass string) {
  // TODO: add conformance classes
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
