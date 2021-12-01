package apif

import(
  "net/http"
  "strings"
  "oaf-server/package/controllers"
  "oaf-server/package/features"
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
  Controller(name string) models.BaseController

  // AddConformanceTemplate(title string, contentType string, renderer coretemplates.RenderConformanceType)
  AddTemplate(name string, title string, contenttype string, rel string, renderer interface{})
  Templates(string, string) []models.Handler

  // Server configuration
  Config() models.Serverconfig

  // Router
  Router() Router
  SetRouter(router Router)

  // Adds a service
  AddService(string, interface{})

  // Returns a service
  GetService(string) interface{}
}

type engine struct {
  // router
  router Router

  // list of controllers
  // path: routename -> controller
  // controllers map[string]apifcontrollers.BaseController
  // list render functions
  // path: routename -> content type -> controller handler function with renderer

  serverconfig models.Serverconfig

  services map[string]interface{}
}

// Returns a new controller for handling OGC Api Feature calls.
//
// This method does not require any furhter configuration
func NewEngine() *engine {
  config := models.NewServerConfig()
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

  landingpageController := &apifcontrollers.LandingpageController{}
  engine.AddRoute(&Routedef{
    Name: "landingpage",
    Path: "",
    Controller: landingpageController,
    LandingpageVisible: true,
  })

  conformanceController := apifcontrollers.NewConformanceController()
  engine.AddRoute(&Routedef{
    Name: "conformance",
    Path: "conformance",
    Controller: conformanceController,
    LandingpageVisible: true,
  })

  return engine
}

func EnableFeatures(engine *engine, service features.FeatureService) {
  engine.AddConformanceClass("http://www.opengis.net/spec/ogcapi-features-1/1.0/conf/features")

  collectionsController := &apifcontrollers.CollectionsController{}
  engine.AddRoute(&Routedef{
    Name: "featurecollections",
    Path: "collections",
    Controller: collectionsController,
    LandingpageVisible: true,
  })
  collectionController := &apifcontrollers.CollectionController{}
  engine.AddRoute(&Routedef{
    Name: "featurecollection",
    Path: "collections/:collection_id",
    Controller: collectionController,
  })
  featuresController := &apifcontrollers.FeaturesController{}
  engine.AddRoute(&Routedef{
    Name: "features",
    Path: "collections/:collection_id/items",
    Controller: featuresController,
  })
  featureController := &apifcontrollers.FeatureController{}
  engine.AddRoute(&Routedef{
    Name: "feature",
    Path: "collections/:collection_id/items/:item_id",
    Controller: featureController,
  })

  engine.AddService("features", service)
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

func (e *engine) AddRoute(routedefinition *Routedef) {
  e.router.AddRoute(routedefinition)
}

func (c *engine) Mount(mountingpath string) {
  mountingpath = sanitizeMountingPath(mountingpath)
  c.Config().SetMountingpath(mountingpath)
}

func (e *engine) Config() models.Serverconfig {
  return e.serverconfig
}

func (e *engine) Templates(url string, contenttype string) []models.Handler {
  result := []models.Handler{}
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

func (e *engine) Controller(name string) models.BaseController {
  return e.router.Controller(name)
}

func (e *engine) AddConformanceClass(conformanceclass string) {
  // TODO: add conformance classes
}

func (e *engine) AddService(name string, service interface{}) {
  e.services[name] = service
}

func (e *engine) GetService(name string) interface{} {
  return e.services[name]
}

func (e *engine) Routes() []models.Route {
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
