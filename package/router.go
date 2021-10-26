package apif

import(
  "net/http"
  "regexp"
  "strings"

  "oaf-server/package/controllers"
  "oaf-server/package/models"
)

type Router interface {
  // Handler for requests
  HandleRequest(w http.ResponseWriter, request *http.Request)

  AddRoute(page string, matchUrl string, controller apifcontrollers.BaseController)

  Route(name string) *Route

  Controller(name string) apifcontrollers.BaseController
  // Controllers() [string]apifcontrollers.BaseController
  //
  // Handlers() [string][string]models.ControllerFunc
}

type router struct {
  routes map[string]*Route
  mountingPath string
}

type Route struct {
  Name string
  MatchUrl string
  Pattern *regexp.Regexp
  Controller apifcontrollers.BaseController
  Handlers map[string]models.ControllerFunc
}

// Initializes a new router object
func NewRouter(mountingPath string) *router {
	router := &router{
    mountingPath: mountingPath,
    routes: make(map[string]*Route),
  }

  return router
}

// Handles requests
func (router *router) HandleRequest(w http.ResponseWriter, r *http.Request) {
  absolutePath := r.URL.EscapedPath()

  // mounting path does not match
  if(!strings.HasPrefix(absolutePath, router.mountingPath)) {
    http.NotFound(w, r)
    return
  }

  // constructing relative path
  mountingPathLength := len(router.mountingPath)
  path := absolutePath[mountingPathLength:]

  // try to match the path
  for _, route := range router.routes {
    if route.Pattern.MatchString(path) {
      contenttype := "application/json"
      handler := route.Handlers[contenttype]
      if handler == nil {
        http.NotFound(w, r)
        return
      }
      
      handler(w, r)
      return
    }
  }

  // end of router
  http.NotFound(w, r)
}

// Adds a new route to the router.
//
// Variables must be named :variable_name, eg: /collections/:collection_id
// --> /collections/(.*)
func (r *router) AddRoute(name string, matchUrl string, controller apifcontrollers.BaseController) {
  pattern := regexp.MustCompile("^" + matchUrl + "$")
  newRoute := &Route{
    Name: name,
    MatchUrl: matchUrl,
    Pattern: pattern,
    Controller: controller,
    Handlers: make(map[string]models.ControllerFunc),
  }

  r.routes[name] = newRoute
}

func (r *router) Route(name string) *Route {
  return r.routes[name]
}

func (r *router) Controller(name string) apifcontrollers.BaseController {
  route := r.Route(name)
  if route == nil {
    return nil
  }

  return route.Controller
}

func (r *Route) AddRoute(contenttype string, handler models.ControllerFunc) {
  r.Handlers[contenttype] = handler
}
