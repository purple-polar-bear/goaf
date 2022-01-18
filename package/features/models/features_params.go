package featuremodels

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
    ContentType:  "application/json",
  }
}
