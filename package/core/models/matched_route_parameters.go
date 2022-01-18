package coremodels

// Parameters of a matched route
// The MatchedRouteParameters contains only the URL matching. It is used
// only for matching against the right controller.
type MatchedRouteParameters interface {
  Get(string) string
}
