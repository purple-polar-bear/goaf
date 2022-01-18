package apicore

import(
  "fmt"
  "net/http"
  "strings"

  "oaf-server/package/core/services"
  "oaf-server/package/core/models"
  "oaf-server/package/core/controllers"
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
  Controller(name string) coremodels.BaseController

  // AddConformanceTemplate(title string, contentType string, renderer coretemplates.RenderConformanceType)
  AddTemplate(name string, title string, contenttype string, rel string, renderer interface{})
  Templates(string, string) []coremodels.Handler

  // Server configuration
  Config() coremodels.Serverconfig

  // Router
  Router() Router
  SetRouter(router Router)

  // Adds a service
  AddService(string, interface{})

  // Returns a service
  GetService(string) interface{}

  // Adds a class to the conformance list
  AddConformanceClass(string)

  // Adds a route to the list of routes
  AddRoute(*Routedef)

  // Returns all the services
  RebuildOpenAPI()
}

type engine struct {
  // router
  router Router

  // list of controllers
  // path: routename -> controller
  // controllers map[string]corecontrollers.BaseController
  // list render functions
  // path: routename -> content type -> controller handler function with renderer

  serverconfig coremodels.Serverconfig

  services map[string]interface{}
}

// Returns a new controller for handling OGC Api Feature calls.
//
// This method does not require any furhter configuration
func NewEngine() *engine {
  config := coremodels.NewServerConfig()
  engine := &engine{
    router: NewRouter(config),
    serverconfig: config,
    services: make(map[string]interface{}),
  }

  return engine
}

func (e *engine) Router() Router {
  return e.router
}

func (e *engine) SetRouter(router Router) {
  e.router = router
}

func NewSimpleEngine(mountingpath string) *engine {
  engine := NewEngine()

  engine.Config().SetMountingpath(mountingpath)

  landingpageController := &corecontrollers.LandingpageController{}
  engine.AddRoute(&Routedef{
    Name: "landingpage",
    Path: "",
    Controller: landingpageController,
    LandingpageVisible: true,
  })

  conformanceController := corecontrollers.NewConformanceController()
  engine.AddRoute(&Routedef{
    Name: "conformance",
    Path: "conformance",
    Controller: conformanceController,
    LandingpageVisible: true,
  })

  apiController := &corecontrollers.APIController{}
  engine.AddRoute(&Routedef{
    Name: "api",
    Path: "api",
    Controller: apiController,
    LandingpageVisible: true,
  })

  service := coreservices.NewCoreService()
  engine.AddService("core", service)

  return engine
}

func EnableAPISpecification(engine *engine) {
}

/*
func EnableTiles(engine *engine) {
  engine.AddConformanceClass("http://www.opengis.net/spec/ogcapi-features-1/1.0/conf/tiles")
  engine.tilesomethingController = corecontrollers.TilesomethingController{}
  router.AddRoute("tiles", "/tiles", engine.tilesomethingController)
}
*/

func (c *engine) HTTPHandler(w http.ResponseWriter, r *http.Request) {
  if (c.router == nil) {
    panic("Apif controller is not mounted")
  }

  coreservice := c.GetService("core").(coreservices.CoreService)
  if(coreservice == nil) {
    panic("Core service is not defined")
  }

  fmt.Printf("Handling request: %s with header %v\n", r.URL.EscapedPath(), r.Form)
  c.router.HandleRequest(coreservice, w, r)
}

func (e *engine) AddRoute(routedefinition *Routedef) {
  e.router.AddRoute(routedefinition)
}

func (c *engine) Mount(mountingpath string) {
  mountingpath = sanitizeMountingPath(mountingpath)
  c.Config().SetMountingpath(mountingpath)
}

func (e *engine) Config() coremodels.Serverconfig {
  return e.serverconfig
}

func (e *engine) Templates(url string, contenttype string) []coremodels.Handler {
  result := []coremodels.Handler{}
  for _, handler := range e.router.Handlers() {
    if url != "" {
      if url != handler.route.name {
        continue
      }
    }

    result = append(result, handler)
  }

  return result
}

func (e *engine) Controller(name string) coremodels.BaseController {
  return e.router.Controller(name)
}

func (e *engine) AddConformanceClass(conformanceclass string) {
  // TODO: add conformance classes
}

func (e *engine) AddService(name string, service interface{}) {
  e.services[name] = service

  configService, ok := service.(coremodels.ConfigurableService)
  if ok {
    configService.SetConfig(e.serverconfig)
  }
}

func (e *engine) GetService(name string) interface{} {
  return e.services[name]
}

func (e *engine) RebuildOpenAPI() {
  service := e.GetService("core").(coreservices.CoreService)
  services := []interface{}{}
  for _, service := range e.services {
    services = append(services, service)
  }

  service.RebuildOpenAPI(services)
}

func (e *engine) Routes() []coremodels.Route {
  return e.router.Routes()
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
