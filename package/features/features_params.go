package features

type FeaturesParams struct {
  CollectionId string
  Limit        int
  Offset       int
  ContentType  string
  Bbox         [4]float64
  Datetime     string
}

func NewFeaturesParams() *FeaturesParams {
  return &FeaturesParams{
    CollectionId: "addresses",
    Limit:        100,
    Offset:       0,
    ContentType:  "application/json",
  }
}
