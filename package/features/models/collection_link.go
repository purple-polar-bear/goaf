package featuremodels

type CollectionLink interface {
  Href() string
  Hreflang() string
  Length() int
  Rel() string
  Title() string
  Type() string
}
