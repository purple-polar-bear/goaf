package viewmodels

type Collections struct {
  Collections []*Collection `json:"collections"`
  Links []*Link `json:"links"`
}

func NewCollections() *Collections {
  return &Collections{
    Collections: []*Collection{},
    Links: []*Link{},
  }
}
