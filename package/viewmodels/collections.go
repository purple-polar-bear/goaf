package viewmodels

import(
  "oaf-server/package/core/viewmodels"
)

type Collections struct {
  Collections []*Collection `json:"collections"`
  Links []*viewmodels.Link `json:"links"`
}

func NewCollections() *Collections {
  return &Collections{
    Collections: []*Collection{},
    Links: []*viewmodels.Link{},
  }
}
