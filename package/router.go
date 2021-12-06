package apif

import(
  "net/http"
  "regexp"
  "strings"

  "oaf-server/package/models"
)

type Router interface {
  // Handler for requests
  HandleRequest(w http.ResponseWriter, request *http.Request)

  AddRoute(routedefinition *Routedef)

  Route(name string) *Route

  Routes() []models.Route
  Controller(name string) models.BaseController
  // Controllers() [string]models.BaseController
  //
  Handlers() []*Handler
}

type router struct {
  routes map[string]*Route
  serverconfig models.Serverconfig
}

type Route struct {
  // Internal name of the route
  name string

  // Raw matchUrl. Variables must be encoded by prefixing with a semicolon
  MatchUrl string

  // Parsed url
  UrlParts []*UrlPart

  // Regular expression
  Pattern *regexp.Regexp

  // Controller to be invoked
  Controller models.BaseController

  // Handlers with templates
  handlers map[string]*Handler

  // Must the controller be visible on the landingpage?
  landingpageVisible  bool
}

type UrlPart struct {
  // name of the url part
  Name string

  // boolean indicating whether this part is a variable or a fixed expression
  IsVariable bool
}

type Handler struct {
  // Route this handler belongs to
  route *Route

  // Contenttype of this handler
  contenttype string

  // title of the handler
  title string

  // relationship of the handler in the API
  rel string

  // actual function
  controllerFunc models.ControllerFunc
}

// Definition struct used for creating new routes
type Routedef struct {
  Name                string
  Path                string
  Controller          models.BaseController
  LandingpageVisible  bool
}

type MatchedRoute struct {
  Parameters map[string]string
}

// Initializes a new router object
func NewRouter(serverconfig models.Serverconfig) *router {
	router := &router{
    serverconfig: serverconfig,
    routes: make(map[string]*Route),
  }

  return router
}

// Handles requests
func (router *router) HandleRequest(w http.ResponseWriter, r *http.Request) {
  absolutePath := r.URL.EscapedPath()
  mountingpath := router.serverconfig.Mountingpath()

  // mounting path does not match
  if(!strings.HasPrefix(absolutePath, mountingpath)) {
    http.NotFound(w, r)
    return
  }

  // constructing relative path
  mountingPathLength := len(mountingpath)
  path := absolutePath[mountingPathLength:]
  pathParts := strings.Split(path, "/")
  pathPartsLen := len(pathParts)
  // remove empty trailing path
  if pathParts[0] == "" {
    pathParts = pathParts[1:]
    pathPartsLen = len(pathParts)
  }
  // remove last empty path for landingpage
  if pathPartsLen == 1 {
    if pathParts[0] == "" {
      pathParts = []string{}
      pathPartsLen = 0
    }
  }

  // try to match the path
  for _, route := range router.routes {
    if len(route.UrlParts) != pathPartsLen {
      continue
    }

    match := true
    routeParameters := NewMatchedRoute()
    for index, part := range route.UrlParts {
      if part.IsVariable {
        routeParameters.Parameters[part.Name] = pathParts[index]
      } else {
        if part.Name != pathParts[index] {
          match = false
        }
      }
    }
    if !match {
      continue
    }

    contenttype := "application/json"
    handler := route.handlers[contenttype]
    if handler == nil {
      http.NotFound(w, r)
      return
    }

    handler.controllerFunc(w, r, routeParameters)
    return
  }

  // end of router
  http.NotFound(w, r)
}

func (r *Route) Name() string {
  return r.name
}

func (r *Route) LandingpageVisible() bool {
  return r.landingpageVisible
}

func (r *Route) Handlers() map[string]models.Handler{
  result := make(map[string]models.Handler)
  for key, value := range r.handlers {
    result[key] = value
  }

  return result
}

// Adds a new route to the router.
//
// Variables must be named :variable_name, eg: /collections/:collection_id
// --> /collections/(.*)
func (r *router) AddRoute(routedefinition *Routedef) {
  name := routedefinition.Name
  matchUrl := routedefinition.Path
  pattern := regexp.MustCompile("^" + matchUrl + "$")
  newRoute := &Route{
    name: name,
    MatchUrl: matchUrl,
    UrlParts: BuildUrlParts(matchUrl),
    landingpageVisible: routedefinition.LandingpageVisible,
    Pattern: pattern,
    Controller: routedefinition.Controller,
    handlers: make(map[string]*Handler),
  }

  r.routes[name] = newRoute
}

func (r *router) Route(name string) *Route {
  return r.routes[name]
}

func (r *router) Routes() []models.Route {
  result := make([]models.Route, 0, len(r.routes))
  for _, route := range r.routes {
    result = append(result, route)
  }

  return result
}

func (r *router) Controller(name string) models.BaseController {
  route := r.Route(name)
  if route == nil {
    return nil
  }

  return route.Controller
}

func (r *Route) AddHandler(contenttype string, handler *Handler) {
  handler.contenttype = contenttype
  handler.route = r
  r.handlers[contenttype] = handler
}

func (r *router) Handlers() []*Handler {
  result := []*Handler{}
  for _, route := range r.routes {
    for _, handler := range route.handlers {
      result = append(result, handler)
    }
  }

  return result
}

func BuildUrlParts(matchUrl string) []*UrlPart {
  result := []*UrlPart{}
  var newUrlPart *UrlPart
  firstElement := true
  for _, element := range strings.Split(matchUrl, "/") {
    if firstElement {
      firstElement = false
      if element == "" {
        continue
      }
    }

    if strings.HasPrefix(element, ":") {
      newUrlPart = &UrlPart{
        Name: element[1:],
        IsVariable: true,
      }
    } else {
      newUrlPart = &UrlPart{
        Name: element,
      }
    }

    result = append(result, newUrlPart)
  }
  return result
}

//
// Handler
//

func (handler *Handler) Title() string {
  return handler.title
}

func (handler *Handler) Rel() string {
  return handler.rel
}

func (handler *Handler) Type() string {
  return handler.contenttype
}

func (handler *Handler) Href(baseUrl string, params map[string]string) string {
  parsedUrl := handler.route.MatchUrl
  for key, value := range params {
    parsedUrl = strings.ReplaceAll(parsedUrl, ":" + key, value)
  }
  return baseUrl + "/" + parsedUrl
}

// Matched route
func NewMatchedRoute() *MatchedRoute {
  return &MatchedRoute{
    Parameters: make(map[string]string),
  }
}

func (route *MatchedRoute) Get(key string) string {
  return route.Parameters[key]
}
