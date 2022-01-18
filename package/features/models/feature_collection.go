package featuremodels

import(
  "oaf-server/package/core/viewmodels"
)

type FeatureCollection struct {
  RequestParams  *FeaturesParams
  Next           bool
  Features       []*Feature
  NumberReturned int64
  Type           string
  Links          []*viewmodels.Link
  NumberMatched  int64
  Crs            string
}

func NewFeatureCollection(params *FeaturesParams, features []*Feature) *FeatureCollection {
  next := (len(features) == params.Limit)
  return &FeatureCollection{
    RequestParams: params,
    Next:          next,
    Features:      features,
  }
}

func (features *FeatureCollection) HasNext() bool {
  return features.Next
}

func (features *FeatureCollection) NextLimit() int {
  if !features.HasNext() {
    return 0
  }

  return features.RequestParams.Limit
}

func (features *FeatureCollection) NextOffset() int {
  return features.RequestParams.Offset + features.RequestParams.Limit
}

func (features *FeatureCollection) Items() []*Feature {
  return features.Features
}
