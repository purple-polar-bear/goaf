package apif

import(
  "net/http"
  "regexp"
  "strings"
)

type Router interface {
  // Handler for requests
  HandleRequest(w http.ResponseWriter, request *http.Request)

  GetRoute(page string)
}

type router struct {
  routes map[string]*route
  mountingPath string
}

type Route struct {
  Page string
  MatchUrl string
  Pattern *regexp.Regexp
  Handler ControllerFunc
}

// Initializes a new router object
func NewRouter(mountingPath string) *router {
	router := &router{
    mountingPath: mountingPath,
  }

  return router
}

// Handles requests
func (router *router) HandleRequest(w http.ResponseWriter, request *http.Request) {
  absolutePath := request.URL.EscapedPath()

  // mounting path does not match
  if(!strings.HasPrefix(absolutePath, router.mountingPath)) {
    http.NotFound(w, request)
    return
  }

  // constructing relative path
  mountingPathLength := len(router.mountingPath)
  path := absolutePath[mountingPathLength:]

  // try to match the path
  for _, route := range router.routes {
    if route.Pattern.MatchString(path) {
      route.Handler(w, request)
      return
    }
  }

  // end of router
  http.NotFound(w, request)
}

// Adds a new route to the router.
//
// Variables must be named :variable_name, eg: /collections/:collection_id
func (r *router) AddRoute(page string, matchUrl string, handler ControllerFunc) {
  pattern := regexp.MustCompile("^" + matchUrl + "$")
  newRoute := &Route{
    Page: page,
    MatchUrl: matchUrl,
    Pattern: pattern,
    Handler: handler,
  }

  r.routes[page] = newRoute
}
