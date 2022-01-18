package featureservices

import (
  "net/http"
  "oaf-server/package/core/services"
  "oaf-server/package/features/models"
)

type FeatureService interface {
  Collections() []featuremodels.Collection
  Collection(string) featuremodels.Collection
  Features(*http.Request, *featuremodels.FeaturesParams) featuremodels.Features
  Feature(string, string) *featuremodels.Feature
  BuildOpenAPISpecification(builder coreservices.OpenAPIBuilder)
}
