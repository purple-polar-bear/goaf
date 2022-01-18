package featuremodels

type Features interface {
  HasNext() bool
  NextLimit() int
  NextOffset() int
  Items() []*Feature
}
