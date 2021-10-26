package models

// A typeroute is a route with a given content type
type Typeroute struct {
  Rel string
  Title string
  Type string
  RelativeHref string
}

func (tr *Typeroute) Href() string {
  return tr.RelativeHref;
}

func (tr *Typeroute) CalculateRelation() string {
  return ""  
}
