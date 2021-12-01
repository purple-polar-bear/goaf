package apifcontrollers

import (
	"net/http"

	"oaf-server/package/features"
	"oaf-server/package/models"
	coretemplates "oaf-server/package/templates/core"
)

type CollectionController struct {
}

func (controller *CollectionController) HandleFunc(app models.Application, r interface{}) models.ControllerFunc {
  renderer := r.(coretemplates.RenderFeaturesType)

  return func(w http.ResponseWriter, r *http.Request) {
    name := "addresses"

    featureService, ok := app.GetService("features").(features.FeatureService)
    if !ok {
      panic("Cannot find featureservice")
    }

    collection := featureService.Collection(name)
    if collection == nil {
      panic("Cannot find collection")
    }

    collectionRoute := app.Templates("featurecollection", "")
    collectionItemsRoute := append(collectionRoute, app.Templates("features", "")...)

    resource := BuildCollection(app, collection, collectionItemsRoute)
    renderer.RenderCollection(models.NewWebcontext(w, r), resource)
  }
}
