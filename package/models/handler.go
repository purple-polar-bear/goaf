package models

// A typeroute is a route with a given content type
type Handler interface {
  Title() string
  Rel() string
  Type() string
  Href(string, map[string]string) string
}

/*
func (handler *Handler) Href(baseUrl string) string {
  return baseUrl + handler.RelativeHref;
}
*/
