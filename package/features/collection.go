package features

type Collection interface {
  Id() string
  Title() string
  Description() string
}
