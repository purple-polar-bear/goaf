package apicore

import(
  "fmt"
  "net/http"
  "regexp"
  "strings"

  "oaf-server/package/core/services"
  "oaf-server/package/core/models"
)

type Router interface {
  // Handler for requests
  HandleRequest(coreservice coreservices.CoreService, w http.ResponseWriter, request *http.Request)

  // Adds a route to the list of routes
  AddRoute(routedefinition *Routedef)

  // Returns a route by name
  Route(name string) *Route

  // Returns an array with all the routes
  Routes() []coremodels.Route

  // Returns a controller
  Controller(name string) coremodels.BaseController

  // Returns all the handlers
  Handlers() []*Handler
}

type router struct {
  routes map[string]*Route
  serverconfig coremodels.Serverconfig
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
  Controller coremodels.BaseController

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
  controllerFunc coremodels.ControllerFunc
}

// Definition struct used for creating new routes
type Routedef struct {
  Name                string
  Path                string
  Controller          coremodels.BaseController
  LandingpageVisible  bool
}

type MatchedRoute struct {
  Parameters map[string]string
}

// Initializes a new router object
func NewRouter(serverconfig coremodels.Serverconfig) *router {
	router := &router{
    serverconfig: serverconfig,
    routes: make(map[string]*Route),
  }

  return router
}

// Handles requests
func (router *router) HandleRequest(coreservice coreservices.CoreService, w http.ResponseWriter, r *http.Request) {
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

  // Parse the forms
  r.ParseForm()

  // construct the accept header
  accept := buildAcceptValue(r, coreservice.ContentTypeUrlEncoder())

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

    handler := findHandler(accept, route.handlers)
    if handler == nil {
      http.NotFound(w, r)
      return
    }

    handler.controllerFunc(handler, w, r, routeParameters)
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

func (r *Route) Handlers() map[string]coremodels.Handler{
  result := make(map[string]coremodels.Handler)
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

func (r *router) Routes() []coremodels.Route {
  result := make([]coremodels.Route, 0, len(r.routes))
  for _, route := range r.routes {
    result = append(result, route)
  }

  return result
}

func (r *router) Controller(name string) coremodels.BaseController {
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

// Builds the relation type value
// It uses the provided content type to determine if 'self' relations must be
// converted to alternate
func (handler *Handler) Rel(contenttype string) string {
  if handler.rel == "self" {
    if handler.contenttype == contenttype {
      return "self"
    }

    return "alternate"
  }

  return handler.rel
}

func (handler *Handler) Type() string {
  return handler.contenttype
}

func (handler *Handler) Href(baseUrl string, params map[string]string, encoder *coremodels.ContentTypeUrlEncoding) string {
  // start with the URL that matched the route
  parsedUrl := handler.route.MatchUrl
  // replace the wildcards with the actual values from the params map
  for key, value := range params {
    parsedUrl = strings.ReplaceAll(parsedUrl, ":" + key, value)
  }

  urlFormat := ""
  if encoder != nil {
    contenttypeEncoding := encoder.ReverseEncodings[handler.contenttype]
    if contenttypeEncoding != "" {
      urlFormat = "?" + encoder.ParameterName + "=" + contenttypeEncoding
    }
  }

  return baseUrl + "/" + parsedUrl + urlFormat
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

func findHandler(accept string, handlers map[string]*Handler) *Handler {
  fmt.Printf("Accept: %v\n", accept)
  for _, contentType := range strings.Split(accept, ",") {
    handler := handlers[contentType]
    if handler != nil {
      return handler
    }
  }

  defaultContentType := "application/json"
  return handlers[defaultContentType]
}

func buildAcceptValue(r *http.Request, contentTypeUrlEncoder *coremodels.ContentTypeUrlEncoding) string {
  acceptHeader := r.Header.Get("accept")
  if contentTypeUrlEncoder == nil {
    return acceptHeader
  }

  urlHeaders := r.Form[contentTypeUrlEncoder.ParameterName]
  if len(urlHeaders) == 0 {
    return acceptHeader
  }

  urlHeader := urlHeaders[0]
  urlHeaderValues := contentTypeUrlEncoder.Encodings[urlHeader]
  if len(urlHeaderValues) == 0 {
    return acceptHeader
  }
  urlHeaderValue := strings.Join(urlHeaderValues, ",")

  acceptList := []string{}
  if contentTypeUrlEncoder.OverrideHeader {
    acceptList = []string{urlHeaderValue, acceptHeader}
  } else {
    acceptList = []string{acceptHeader, urlHeaderValue}
  }
  fmt.Printf("%v\n", acceptList)

  return strings.Join(acceptList, ",")
}
