package apifcontrollers

import(
  "net/http"

  "oaf-server/package/core"
	"oaf-server/package/features"
  "oaf-server/package/models"
	coretemplates "oaf-server/package/templates/core"
)

type CollectionController struct {
}

func (controller *CollectionController) HandleFunc(app models.Application, r interface{}) models.ControllerFunc {
  renderer := r.(coretemplates.RenderFeaturesType)

  return func(handler models.Handler, w http.ResponseWriter, r *http.Request, routeParameters models.MatchedRouteParameters) {
    name := routeParameters.Get("collection_id")

    featureService, ok := app.GetService("features").(features.FeatureService)
    if !ok {
      panic("Cannot find featureservice")
    }

    coreservice, ok := app.GetService("core").(apifcore.CoreService)
    if !ok {
      panic("Cannot find coreservice")
    }

    encoding := coreservice.ContentTypeUrlEncoder()

    collection := featureService.Collection(name)
    if collection == nil {
      // panic("Cannot find collection")
      http.NotFound(w, r)
      return
    }

    collectionRoute := app.Templates("featurecollection", "")
    collectionItemsRoute := append(collectionRoute, app.Templates("features", "")...)

    resource := BuildCollection(handler, app, encoding, collection, collectionItemsRoute)
    renderer.RenderCollection(models.NewWebcontext(w, r), resource)
  }
}
