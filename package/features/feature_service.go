package features

type FeatureService interface {
  Collections() []Collection
  Collection(string) Collection
  Features(*FeaturesParams) Features
}
