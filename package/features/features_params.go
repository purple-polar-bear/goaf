package features

type FeaturesParams struct {
  CollectionId string
  Limit int
  Offset int
  ContentType string
  Bbox string
  Datetime string
}

func NewFeaturesParams() *FeaturesParams {
  return &FeaturesParams{
    Limit: 100,
    Offset: 0,
    ContentType: "application/json",
  }
}
