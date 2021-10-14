package models

type Engine interface {
  Title() string
  Description() string
  Templates() []templates
}
