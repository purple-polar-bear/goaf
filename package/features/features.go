package features

type Features interface {
  HasNext() bool
  NextLimit() int
  NextOffset() int
  Items() []*Feature
}

// Implementation
type SimpleFeatures struct {
  RequestParams *FeaturesParams
  Next bool
  Features []*Feature
}

func NewSimpleFeatures(params *FeaturesParams, features []*Feature) *SimpleFeatures {
  next := (len(features) == params.Limit)
  return &SimpleFeatures{
    RequestParams: params,
    Next: next,
    Features: features,
  }
}

func (features *SimpleFeatures) HasNext() bool {
  return features.Next
}

func (features *SimpleFeatures) NextLimit() int {
  if !features.HasNext() {
    return 0
  }

  return features.RequestParams.Limit
}

func (features *SimpleFeatures) NextOffset() int {
  return features.RequestParams.Offset + features.RequestParams.Limit
}

func (features *SimpleFeatures) Items() []*Feature {
  return features.Features
}
