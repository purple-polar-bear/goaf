package features

import "net/http"

type FeatureService interface {
  Collections() []Collection
  Collection(string) Collection
  Features(*http.Request, *FeaturesParams) Features
}
