package coremodels

// A typeroute is a route with a given content type
type Handler interface {
  Title() string

  // Determine the relation-type of this handler. The parameter contains the
  // actual contenttype
  Rel(string) string

  Type() string

  // Calculates the full URL of the handler
  Href(string, map[string]string, *ContentTypeUrlEncoding) string
}
