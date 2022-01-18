package featurescontrollers

import(
  "net/http"

  "oaf-server/package/core/services"
  "oaf-server/package/core/models"
  "oaf-server/package/features/services"
  "oaf-server/package/features/templates/core"
)

type CollectionController struct {
}

func (controller *CollectionController) HandleFunc(app coremodels.Application, r interface{}) coremodels.ControllerFunc {
  renderer := r.(coretemplates.RenderFeaturesType)

  return func(handler coremodels.Handler, w http.ResponseWriter, r *http.Request, routeParameters coremodels.MatchedRouteParameters) {
    name := routeParameters.Get("collection_id")

    featureService, ok := app.GetService("features").(featureservices.FeatureService)
    if !ok {
      panic("Cannot find featureservice")
    }

    coreservice, ok := app.GetService("core").(coreservices.CoreService)
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
    renderer.RenderCollection(coremodels.NewWebcontext(w, r), resource)
  }
}
