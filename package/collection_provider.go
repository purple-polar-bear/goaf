package apif

import(
  "oaf-server/package/controllers"
)

type CollectionProvider interface {
  All() []*CollectionInfo

  Get(name string) *CollectionInfo
}

type CollectionInfo struct {
  Name string
}

type FeatureApplication struct {
  CollectionsController apifcontrollers.CollectionsController
  CollectionController apifcontrollers.CollectionController

  providers []CollectionProvider
}

func (app *FeatureApplication) AddCollectionProvider(provider CollectionProvider) {
  app.providers = append(app.providers, provider)
}

func (app *FeatureApplication) All() []*CollectionInfo {
  result := []*CollectionInfo{}
  for _, provider := range app.providers {
    result = append(result, provider.All()...)
  }

  return result
}

func (app *FeatureApplication) Get(name string) *CollectionInfo {
  for _, provider := range app.providers {
    candidate := provider.Get(name)
    if candidate != nil {
      return candidate
    }
  }

  return nil
}
