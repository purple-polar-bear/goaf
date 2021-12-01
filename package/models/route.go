package models

type Route interface {
  // Name of the route
  Name() string

  // List of handlers. The indexing string represents the contenttype
  Handlers() map[string]Handler

  LandingpageVisible()  bool
}

// Parameters of a matched route
type MatchedRouteParameters interface {
  Get(string) string
}
