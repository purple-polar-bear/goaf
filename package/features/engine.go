package apifeatures

import(
  "oaf-server/package/core"
  "oaf-server/package/features/controllers"
  "oaf-server/package/features/services"
)

func EnableFeatures(engine apicore.Engine, service featureservices.FeatureService) {
  engine.AddConformanceClass("http://www.opengis.net/spec/ogcapi-features-1/1.0/conf/features")

  collectionsController := &featurescontrollers.CollectionsController{}
  engine.AddRoute(&apicore.Routedef{
    Name: "featurecollections",
    Path: "collections",
    Controller: collectionsController,
    LandingpageVisible: true,
  })
  collectionController := &featurescontrollers.CollectionController{}
  engine.AddRoute(&apicore.Routedef{
    Name: "featurecollection",
    Path: "collections/:collection_id",
    Controller: collectionController,
  })
  featuresController := &featurescontrollers.FeaturesController{}
  engine.AddRoute(&apicore.Routedef{
    Name: "features",
    Path: "collections/:collection_id/items",
    Controller: featuresController,
  })
  featureController := &featurescontrollers.FeatureController{}
  engine.AddRoute(&apicore.Routedef{
    Name: "feature",
    Path: "collections/:collection_id/items/:item_id",
    Controller: featureController,
  })

  engine.AddService("features", service)
}
